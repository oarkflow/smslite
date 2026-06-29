package smslite

import "strings"

type NormalizationResult struct {
	Original    string       `json:"original"`
	Normalized  string       `json:"normalized"`
	Changed     bool         `json:"changed"`
	Suggestions []Suggestion `json:"suggestions"`
}

func NormalizeSmartCharacters(text string) NormalizationResult { return NormalizeForGSM7(text) }

func NormalizeForGSM7(text string) NormalizationResult {
	sugs := SuggestReplacements(text)
	if len(sugs) == 0 {
		return NormalizationResult{Original: text, Normalized: text}
	}
	var b strings.Builder
	b.Grow(len(text))
	ri := 0
	for _, r := range text {
		repl, ok := replacementFor(r)
		if ok {
			b.WriteString(repl)
		} else {
			b.WriteRune(r)
		}
		ri++
	}
	return NormalizationResult{Original: text, Normalized: b.String(), Changed: true, Suggestions: sugs}
}

func CanReduceToGSM7(text string) bool { return IsGSM7(NormalizeForGSM7(text).Normalized) }

func SuggestGSM7Replacement(text string) []Suggestion { return SuggestReplacements(text) }

func SuggestReplacements(text string) []Suggestion {
	var out []Suggestion
	ri := 0
	for _, r := range text {
		if repl, ok := replacementFor(r); ok {
			saves := UCS2UnitLength(string(r))
			if c, ok := GSM7SeptetLength(repl); ok == nil {
				saves -= c
			}
			out = append(out, Suggestion{IndexRune: ri, Original: string(r), Replacement: repl, Reason: "replace Unicode punctuation/invisible character with GSM-7 equivalent", SavesUnits: saves})
		}
		ri++
	}
	return out
}

func replacementFor(r rune) (string, bool) {
	switch r {
	case '“', '”':
		return "\"", true
	case '‘', '’', '`':
		return "'", true
	case '–', '—', '−':
		return "-", true
	case '…':
		return "...", true
	case 0x00A0:
		return " ", true
	case 0x200B, 0x200C, 0x200D, 0x2060, 0xFEFF:
		return "", true
	default:
		return "", false
	}
}
