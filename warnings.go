package smslite

import "fmt"

func BuildWarnings(text string, chars []CharacterInfo, seg Segmentation) []Warning {
	var out []Warning
	if seg.IsMultipart {
		out = append(out, Warning{Code: "MULTIPART", Message: fmt.Sprintf("message uses %d SMS parts", seg.SegmentCount), Suggestion: "shorten text to reduce billable parts"})
	}
	for _, c := range chars {
		if c.IsUCS2 {
			out = append(out, Warning{Code: "UCS2_FORCED", Message: "character forces UCS-2 encoding", IndexRune: c.IndexRune, Character: c.Char, Suggestion: "remove or replace this character to keep GSM-7"})
		}
		if c.IsGSMExtended {
			out = append(out, Warning{Code: "GSM_EXTENSION_COST", Message: "GSM extension character costs 2 septets", IndexRune: c.IndexRune, Character: c.Char})
		}
		if c.IsInvisible && c.Char != " " && c.Char != "\n" && c.Char != "\r" {
			code := "INVISIBLE_CHARACTER"
			suggestion := "remove it unless intentional"
			switch c.InvisibleClass {
			case InvisibleZeroWidth:
				code = "ZERO_WIDTH_CHARACTER"
			case InvisibleBidiControl:
				code = "BIDI_CONTROL_CHARACTER"
				suggestion = "remove it unless bidirectional text control is intentional; it can spoof visual order"
			case InvisibleFormatControl:
				code = "FORMAT_CONTROL_CHARACTER"
			case InvisibleUnicodeSpace:
				code = "UNICODE_SPACE_CHARACTER"
				suggestion = "replace with regular ASCII space when possible"
			case InvisibleMathInvisible:
				code = "MATH_INVISIBLE_CHARACTER"
			case InvisibleVariationSelector:
				code = "VARIATION_SELECTOR"
			}
			out = append(out, Warning{Code: code, Message: c.InvisibleClassLabel + " found: " + c.Visible, IndexRune: c.IndexRune, Character: c.Char, Suggestion: suggestion})
		}
		if c.IsEmoji {
			out = append(out, Warning{Code: "EMOJI", Message: "emoji requires UCS-2 and may use two UTF-16 units", IndexRune: c.IndexRune, Character: c.Char})
		}
		if c.IsCombiningMark {
			out = append(out, Warning{Code: "COMBINING_MARK", Message: "combining mark affects visual length", IndexRune: c.IndexRune, Character: c.Char})
		}
	}
	return out
}
