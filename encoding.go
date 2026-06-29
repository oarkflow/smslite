package smslite

type EncodingResult struct {
	Encoding Encoding `json:"encoding"`
	Reason   string   `json:"reason"`
	AtRune   int      `json:"atRune,omitempty"`
	AtByte   int      `json:"atByte,omitempty"`
	Char     string   `json:"char,omitempty"`
}

// DetectEncoding is a zero-allocation hot-path detector. For details about the
// first UCS-2-forcing character, use DetectEncodingDetailed or FirstUCS2CharacterPosition.
func DetectEncoding(text string) Encoding { return AnalyzeSummary(text).Encoding }

func DetectEncodingDetailed(text string) EncodingResult {
	ri := 0
	for i := 0; i < len(text); {
		b := text[i]
		if b < 0x80 {
			if asciiGSM7Cost[b] == 0 {
				return EncodingResult{Encoding: EncodingUCS2, Reason: "contains character outside GSM-7 alphabet", AtRune: ri, AtByte: i, Char: string(rune(b))}
			}
			ri++
			i++
			continue
		}
		r, sz := decodeRuneAt(text, i)
		if _, ok := GSM7CharCost(r); !ok {
			return EncodingResult{Encoding: EncodingUCS2, Reason: "contains character outside GSM-7 alphabet", AtRune: ri, AtByte: i, Char: string(r)}
		}
		ri++
		i += sz
	}
	return EncodingResult{Encoding: EncodingGSM7, Reason: "all characters fit GSM-7 default or extension alphabet"}
}
