package smslite

// Summary is the allocation-free analysis result for hot paths. It excludes slices,
// rendered strings, and per-character reports. Use Analyze for full diagnostics.
type Summary struct {
	Encoding           Encoding `json:"encoding"`
	ByteLength         int      `json:"byteLength"`
	RuneCount          int      `json:"runeCount"`
	GSMSeptetCount     int      `json:"gsmSeptetCount"`
	UCS2UnitCount      int      `json:"ucs2UnitCount"`
	LengthUnits        int      `json:"lengthUnits"`
	SegmentCount       int      `json:"segmentCount"`
	IsMultipart        bool     `json:"isMultipart"`
	SingleSegmentLimit int      `json:"singleSegmentLimit"`
	MultipartLimit     int      `json:"multipartLimit"`
	FirstUCS2RuneIndex int      `json:"firstUCS2RuneIndex"`
	FirstUCS2ByteIndex int      `json:"firstUCS2ByteIndex"`
}

// AnalyzeSummary performs allocation-free SMS encoding and segment analysis.
func AnalyzeSummary(text string) Summary { return AnalyzeSummaryASCII(text) }

func analyzeSummaryGeneric(text string) Summary {
	gsmUnits, ucsUnits, runes := 0, 0, 0
	firstUCS2Rune, firstUCS2Byte := -1, -1
	gsmOK := true

	for i := 0; i < len(text); {
		b := text[i]
		if b < 0x80 {
			if c := asciiGSM7Cost[b]; c != 0 && gsmOK {
				gsmUnits += int(c)
			} else if gsmOK {
				firstUCS2Rune, firstUCS2Byte = runes, i
				gsmOK = false
			}
			ucsUnits++
			runes++
			i++
			continue
		}

		r, sz := decodeRuneAt(text, i)
		if c, ok := GSM7CharCost(r); ok && gsmOK {
			gsmUnits += c
		} else if gsmOK {
			firstUCS2Rune, firstUCS2Byte = runes, i
			gsmOK = false
		}
		if r < 0x10000 {
			ucsUnits++
		} else {
			ucsUnits += 2
		}
		runes++
		i += sz
	}

	enc := EncodingGSM7
	units, single, multi := gsmUnits, GSM7SingleLimit, GSM7MultipartLimit
	if !gsmOK {
		enc, units, single, multi = EncodingUCS2, ucsUnits, UCS2SingleLimit, UCS2MultipartLimit
	}
	parts := 0
	if units > 0 {
		if units <= single {
			parts = 1
		} else {
			parts = (units + multi - 1) / multi
		}
	}
	return Summary{Encoding: enc, ByteLength: len(text), RuneCount: runes, GSMSeptetCount: gsmUnits, UCS2UnitCount: ucsUnits, LengthUnits: units, SegmentCount: parts, IsMultipart: parts > 1, SingleSegmentLimit: single, MultipartLimit: multi, FirstUCS2RuneIndex: firstUCS2Rune, FirstUCS2ByteIndex: firstUCS2Byte}
}

// AnalyzeSummaryASCII is the fastest hot path for already-known ASCII payloads.
// It is allocation-free and avoids UTF-8 decoding. Non-ASCII bytes are treated as UCS-2 forcing.
func AnalyzeSummaryASCII(text string) Summary {
	gsmUnits := 0
	first := -1
	for i := 0; i < len(text); i++ {
		c := asciiGSM7Cost[text[i]&0x7f]
		if text[i] >= 0x80 || c == 0 {
			first = i
			break
		}
		gsmUnits += int(c)
	}
	if first >= 0 {
		// Non-ASCII or non-GSM ASCII was found; delegate to the full zero-alloc scanner.
		return analyzeSummaryGeneric(text)
	}
	parts := 0
	if gsmUnits > 0 {
		if gsmUnits <= GSM7SingleLimit {
			parts = 1
		} else {
			parts = (gsmUnits + GSM7MultipartLimit - 1) / GSM7MultipartLimit
		}
	}
	return Summary{Encoding: EncodingGSM7, ByteLength: len(text), RuneCount: len(text), GSMSeptetCount: gsmUnits, UCS2UnitCount: len(text), LengthUnits: gsmUnits, SegmentCount: parts, IsMultipart: parts > 1, SingleSegmentLimit: GSM7SingleLimit, MultipartLimit: GSM7MultipartLimit, FirstUCS2RuneIndex: -1, FirstUCS2ByteIndex: -1}
}
