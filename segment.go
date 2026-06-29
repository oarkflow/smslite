package smslite

import "unicode/utf8"

func CalculateSegments(text string) Segmentation {
	return CalculateSegmentsWithOptions(text, DefaultOptions())
}

func CalculateSegmentsDetailed(text string) Segmentation {
	return CalculateSegmentsWithOptions(text, DiagnosticOptions())
}

func CalculateSegmentsWithOptions(text string, opt Options) Segmentation {
	encRes := DetectEncodingDetailed(text)
	enc := encRes.Encoding
	units := 0
	single, multi := GSM7SingleLimit, GSM7MultipartLimit
	if enc == EncodingGSM7 {
		units, _ = GSM7SeptetLength(text)
	} else {
		units = UCS2UnitLength(text)
		single, multi = UCS2SingleLimit, UCS2MultipartLimit
	}
	cnt := 0
	limit := single
	if units == 0 {
		cnt = 0
	} else if units <= single {
		cnt = 1
	} else {
		cnt = (units + multi - 1) / multi
		limit = multi
	}
	seg := Segmentation{Encoding: enc, LengthUnits: units, SegmentCount: cnt, IsMultipart: cnt > 1, SingleSegmentLimit: single, MultipartLimit: multi, PerSegmentLimit: limit}
	if cnt > 1 {
		u := BuildUDH(opt, cnt, 1)
		seg.UDH = &u
	}
	if opt.CollectSegments {
		seg.Segments, _ = SplitSMSWithOptions(text, opt)
	}
	return seg
}

func SplitSMS(text string) ([]SMSSegment, error) { return SplitSMSWithOptions(text, DefaultOptions()) }

func SplitSMSWithOptions(text string, opt Options) ([]SMSSegment, error) {
	if DetectEncoding(text) == EncodingGSM7 {
		return SplitGSM7WithOptions(text, opt)
	}
	return SplitUCS2WithOptions(text, opt)
}

func SplitGSM7(text string) ([]SMSSegment, error) {
	return SplitGSM7WithOptions(text, DefaultOptions())
}
func SplitUCS2(text string) ([]SMSSegment, error) {
	return SplitUCS2WithOptions(text, DefaultOptions())
}

func SplitGSM7WithOptions(text string, opt Options) ([]SMSSegment, error) {
	units, ok := GSM7SeptetLengthFast(text)
	if !ok {
		return nil, ErrNonGSM7
	}
	return splitByUnits(text, EncodingGSM7, units, opt)
}

func SplitUCS2WithOptions(text string, opt Options) ([]SMSSegment, error) {
	return splitByUnits(text, EncodingUCS2, UCS2UnitLength(text), opt)
}

func splitByUnits(text string, enc Encoding, totalUnits int, opt Options) ([]SMSSegment, error) {
	if totalUnits == 0 {
		return nil, nil
	}
	single, multi := GSM7SingleLimit, GSM7MultipartLimit
	if enc == EncodingUCS2 {
		single, multi = UCS2SingleLimit, UCS2MultipartLimit
	}
	limit := single
	totalParts := 1
	if totalUnits > single {
		limit = multi
		totalParts = (totalUnits + multi - 1) / multi
	}
	if opt.MaxParts > 0 && totalParts > opt.MaxParts {
		return nil, ErrTooManyParts
	}
	segments := make([]SMSSegment, 0, totalParts)
	startByte, startRune, curUnits, ri := 0, 0, 0, 0
	lastBreakByte, lastBreakRune, lastBreakUnits := -1, -1, 0
	for bi, r := range text {
		cost := UCS2Units(r)
		if enc == EncodingGSM7 {
			cost, _ = GSM7CharCost(r)
		}
		if opt.PreserveWords && (r == ' ' || r == '\n' || r == '\t') {
			lastBreakByte, lastBreakRune, lastBreakUnits = bi+utf8.RuneLen(r), ri+1, curUnits+cost
		}
		if curUnits+cost > limit {
			endByte, endRune, endUnits := bi, ri, curUnits
			if opt.PreserveWords && lastBreakByte > startByte && lastBreakUnits > 0 {
				endByte, endRune, endUnits = lastBreakByte, lastBreakRune, lastBreakUnits
			}
			segments = append(segments, makeSegment(text, enc, len(segments), totalParts, startByte, endByte, startRune, endRune, endUnits, opt))
			if endByte == bi {
				startByte, startRune, curUnits = bi, ri, cost
			} else {
				startByte, startRune, curUnits = endByte, endRune, 0
				lastBreakByte, lastBreakRune, lastBreakUnits = -1, -1, 0
				continue
			}
		} else {
			curUnits += cost
		}
		ri++
	}
	if startByte < len(text) || curUnits > 0 {
		segments = append(segments, makeSegment(text, enc, len(segments), totalParts, startByte, len(text), startRune, ri, curUnits, opt))
	}
	// recompute total after preserve-word continuation edge
	for i := range segments {
		segments[i].TotalParts = len(segments)
		segments[i].PartNumber = i + 1
		if len(segments) > 1 {
			u := BuildUDH(opt, len(segments), i+1)
			segments[i].UDH = &u
		}
	}
	return segments, nil
}

func makeSegment(text string, enc Encoding, idx, total, sb, eb, sr, er, units int, opt Options) SMSSegment {
	s := SMSSegment{Index: idx, PartNumber: idx + 1, TotalParts: total, Text: text[sb:eb], LengthUnits: units, StartRuneIndex: sr, EndRuneIndex: er, StartByteIndex: sb, EndByteIndex: eb}
	if total > 1 {
		u := BuildUDH(opt, total, idx+1)
		s.UDH = &u
	}
	if opt.CollectCharacters {
		s.Characters = BuildCharacterInfo(s.Text)
		for i := range s.Characters {
			s.Characters[i].SegmentIndex = idx
			s.Characters[i].PositionInSegment = i
		}
	}
	return s
}
