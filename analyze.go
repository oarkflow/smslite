package smslite

const (
	reasonGSM7 = "all characters fit GSM-7 default or extension alphabet"
	reasonUCS2 = "contains character outside GSM-7 alphabet"
)

// Analyze is the allocation-free hot-path SMS analyzer. It returns complete
// message-level facts: encoding, length units, segment count, multipart limits,
// and first UCS-2-forcing position. It intentionally does not allocate or fill
// per-character slices, rendered reports, warnings, or segment slices.
//
// For complete diagnostic reports use AnalyzeDetailed, or reuse an Analyzer with
// AnalyzeInto to amortize allocations across calls.
func Analyze(text string) Summary { return AnalyzeSummary(text) }

// AnalyzeFast is an explicit alias for Analyze.
func AnalyzeFast(text string) Summary { return Analyze(text) }

// AnalyzeWithOptions preserves the allocation-free behavior when no collection
// options are enabled. If any collection flag is enabled it delegates to
// AnalyzeDetailedWithOptions.
func AnalyzeWithOptions(text string, opt Options) SMSAnalysis {
	if opt.CollectCharacters || opt.CollectWarnings || opt.CollectSegments || opt.CollectVisualization {
		return AnalyzeDetailedWithOptions(text, opt)
	}
	return analysisFromSummary(text, AnalyzeSummary(text))
}

func analysisFromSummary(text string, s Summary) SMSAnalysis {
	reason := reasonGSM7
	if s.Encoding == EncodingUCS2 {
		reason = reasonUCS2
	}
	max := s.SingleSegmentLimit
	if s.IsMultipart {
		max = s.MultipartLimit
	}
	return SMSAnalysis{Text: text, Encoding: s.Encoding, EncodingReason: reason, ByteLength: s.ByteLength, RuneCount: s.RuneCount, GSMSeptetCount: s.GSMSeptetCount, UCS2UnitCount: s.UCS2UnitCount, LengthUnits: s.LengthUnits, SegmentCount: s.SegmentCount, MaxCharsPerSegment: max, SingleSegmentLimit: s.SingleSegmentLimit, MultipartLimit: s.MultipartLimit, IsMultipart: s.IsMultipart, FirstUCS2RuneIndex: s.FirstUCS2RuneIndex, FirstUCS2ByteIndex: s.FirstUCS2ByteIndex}
}

// AnalyzeDetailed builds the full diagnostic report: characters, UCS-2 forcing
// characters, invisible characters, warnings, recommendations, and segments.
// This allocates by design because it returns slices and rendered strings.
func AnalyzeDetailed(text string) SMSAnalysis {
	return AnalyzeDetailedWithOptions(text, DiagnosticOptions())
}

func AnalyzeDetailedWithOptions(text string, opt Options) SMSAnalysis {
	opt.CollectCharacters = true
	opt.CollectWarnings = true
	opt.CollectSegments = true
	enc := DetectEncodingDetailed(text)
	gsmLen, gsmOK := GSM7SeptetLengthFast(text)
	ucsLen := UCS2UnitLength(text)
	seg := CalculateSegmentsWithOptions(text, opt)
	chars := BuildCharacterInfo(text)
	assignCharacterSegments(chars, seg.Segments)
	ucsChars := AppendUCS2Characters(nil, text)
	invisChars := AppendInvisibleCharacterPositions(nil, text)
	warnings := BuildWarnings(text, chars, seg)
	sugs := SuggestReplacements(text)
	lengthUnits := ucsLen
	if enc.Encoding == EncodingGSM7 && gsmOK {
		lengthUnits = gsmLen
	}
	max := seg.SingleSegmentLimit
	if seg.IsMultipart {
		max = seg.MultipartLimit
	}
	return SMSAnalysis{Text: text, Encoding: enc.Encoding, EncodingReason: enc.Reason, ByteLength: len(text), RuneCount: RuneCount(text), GSMSeptetCount: gsmLen, UCS2UnitCount: ucsLen, LengthUnits: lengthUnits, SegmentCount: seg.SegmentCount, MaxCharsPerSegment: max, SingleSegmentLimit: seg.SingleSegmentLimit, MultipartLimit: seg.MultipartLimit, IsMultipart: seg.IsMultipart, Segments: seg.Segments, Characters: chars, UCS2Characters: ucsChars, InvisibleCharacters: invisChars, Warnings: warnings, Recommendations: sugs, FirstUCS2RuneIndex: AnalyzeSummary(text).FirstUCS2RuneIndex, FirstUCS2ByteIndex: AnalyzeSummary(text).FirstUCS2ByteIndex}
}

func assignCharacterSegments(chars []CharacterInfo, segs []SMSSegment) {
	if len(chars) == 0 || len(segs) == 0 {
		return
	}
	si := 0
	for i := range chars {
		for si+1 < len(segs) && chars[i].IndexByte >= segs[si].EndByteIndex {
			si++
		}
		chars[i].SegmentIndex = si
		chars[i].PositionInSegment = chars[i].IndexRune - segs[si].StartRuneIndex
	}
}

func EstimateBillingUnits(text string) BillingInfo {
	s := AnalyzeSummary(text)
	reason := "single part"
	if s.IsMultipart {
		reason = "multipart message"
	}
	return BillingInfo{Encoding: s.Encoding, SegmentCount: s.SegmentCount, BillableParts: s.SegmentCount, Reason: reason}
}

func IsBillableMultipart(text string) bool { return AnalyzeSummary(text).IsMultipart }

func RemainingUnits(text string) int {
	s := AnalyzeSummary(text)
	if s.SegmentCount == 0 {
		return s.SingleSegmentLimit
	}
	limit := s.SingleSegmentLimit
	if s.IsMultipart {
		limit = s.MultipartLimit
	}
	rem := limit - (s.LengthUnits % limit)
	if rem == limit {
		return 0
	}
	return rem
}

func RemainingCharacters(text string) int { return RemainingUnits(text) }

func NextCharacterCost(r rune, currentText string) int {
	if DetectEncoding(currentText) == EncodingGSM7 {
		if c, ok := GSM7CharCost(r); ok {
			return c
		}
		return UCS2UnitLength(currentText+string(r)) - UCS2UnitLength(currentText)
	}
	return UCS2Units(r)
}
