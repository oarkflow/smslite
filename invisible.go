package smslite

import "fmt"

// InvisibleClass describes why a rune is visually confusing in SMS/UI text.
type InvisibleClass uint8

const (
	InvisibleNone InvisibleClass = iota
	InvisibleASCIIControl
	InvisibleWhitespace
	InvisibleUnicodeSpace
	InvisibleZeroWidth
	InvisibleFormatControl
	InvisibleBidiControl
	InvisibleVariationSelector
	InvisibleCombiningMark
	InvisibleMathInvisible
)

func (c InvisibleClass) String() string {
	switch c {
	case InvisibleASCIIControl:
		return "ASCII control"
	case InvisibleWhitespace:
		return "whitespace"
	case InvisibleUnicodeSpace:
		return "Unicode space separator"
	case InvisibleZeroWidth:
		return "zero-width character"
	case InvisibleFormatControl:
		return "format control"
	case InvisibleBidiControl:
		return "bidirectional control"
	case InvisibleVariationSelector:
		return "variation selector"
	case InvisibleCombiningMark:
		return "combining mark"
	case InvisibleMathInvisible:
		return "mathematical invisible operator"
	default:
		return "visible"
	}
}

// InvisibleCharacterPosition is a compact record for invisible, zero-width,
// bidi, format-control, spacing, and combining characters that can alter SMS
// rendering, billing, or user perception.
type InvisibleCharacterPosition struct {
	IndexByte  int            `json:"indexByte"`
	IndexRune  int            `json:"indexRune"`
	IndexUTF16 int            `json:"indexUTF16"`
	Char       string         `json:"char"`
	Visible    string         `json:"visible"`
	CodePoint  string         `json:"codePoint"`
	Hex        string         `json:"hex"`
	Name       string         `json:"name"`
	Class      InvisibleClass `json:"class"`
	ClassLabel string         `json:"classLabel"`
	Reason     string         `json:"reason"`
}

func InvisibleInfo(r rune) (InvisibleClass, string, string, bool) {
	if r < 0x20 || r == 0x7F {
		return InvisibleASCIIControl, UnicodeName(r), fmt.Sprintf("control character %s is not normally visible", CodePoint(r)), true
	}
	switch r {
	case ' ', '\t', '\n', '\r', '\v', '\f':
		return InvisibleWhitespace, UnicodeName(r), "ASCII whitespace affects layout but is easy to miss visually", true
	case 0x00A0:
		return InvisibleUnicodeSpace, "NO-BREAK SPACE", "non-breaking space looks like a normal space but is not GSM-7", true
	case 0x1680:
		return InvisibleUnicodeSpace, "OGHAM SPACE MARK", "Unicode space separator looks blank and forces UCS-2", true
	case 0x180E:
		return InvisibleFormatControl, "MONGOLIAN VOWEL SEPARATOR", "deprecated invisible format character", true
	case 0x2000:
		return InvisibleUnicodeSpace, "EN QUAD", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2001:
		return InvisibleUnicodeSpace, "EM QUAD", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2002:
		return InvisibleUnicodeSpace, "EN SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2003:
		return InvisibleUnicodeSpace, "EM SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2004:
		return InvisibleUnicodeSpace, "THREE-PER-EM SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2005:
		return InvisibleUnicodeSpace, "FOUR-PER-EM SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2006:
		return InvisibleUnicodeSpace, "SIX-PER-EM SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2007:
		return InvisibleUnicodeSpace, "FIGURE SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2008:
		return InvisibleUnicodeSpace, "PUNCTUATION SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2009:
		return InvisibleUnicodeSpace, "THIN SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x200A:
		return InvisibleUnicodeSpace, "HAIR SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x200B:
		return InvisibleZeroWidth, "ZERO WIDTH SPACE", "zero-width space is invisible and forces UCS-2", true
	case 0x200C:
		return InvisibleZeroWidth, "ZERO WIDTH NON-JOINER", "zero-width non-joiner is invisible and forces UCS-2", true
	case 0x200D:
		return InvisibleZeroWidth, "ZERO WIDTH JOINER", "zero-width joiner is invisible and forces UCS-2", true
	case 0x200E:
		return InvisibleBidiControl, "LEFT-TO-RIGHT MARK", "bidirectional mark is invisible and can change display order", true
	case 0x200F:
		return InvisibleBidiControl, "RIGHT-TO-LEFT MARK", "bidirectional mark is invisible and can change display order", true
	case 0x2028:
		return InvisibleUnicodeSpace, "LINE SEPARATOR", "Unicode line separator is visually disruptive and forces UCS-2", true
	case 0x2029:
		return InvisibleUnicodeSpace, "PARAGRAPH SEPARATOR", "Unicode paragraph separator is visually disruptive and forces UCS-2", true
	case 0x202A:
		return InvisibleBidiControl, "LEFT-TO-RIGHT EMBEDDING", "bidirectional embedding control is invisible and can spoof visual order", true
	case 0x202B:
		return InvisibleBidiControl, "RIGHT-TO-LEFT EMBEDDING", "bidirectional embedding control is invisible and can spoof visual order", true
	case 0x202C:
		return InvisibleBidiControl, "POP DIRECTIONAL FORMATTING", "bidirectional formatting control is invisible", true
	case 0x202D:
		return InvisibleBidiControl, "LEFT-TO-RIGHT OVERRIDE", "bidirectional override is invisible and can spoof visual order", true
	case 0x202E:
		return InvisibleBidiControl, "RIGHT-TO-LEFT OVERRIDE", "bidirectional override is invisible and can spoof visual order", true
	case 0x202F:
		return InvisibleUnicodeSpace, "NARROW NO-BREAK SPACE", "narrow no-break space looks blank and forces UCS-2", true
	case 0x205F:
		return InvisibleUnicodeSpace, "MEDIUM MATHEMATICAL SPACE", "Unicode spacing character looks blank and forces UCS-2", true
	case 0x2060:
		return InvisibleZeroWidth, "WORD JOINER", "word joiner is invisible and forces UCS-2", true
	case 0x2061:
		return InvisibleMathInvisible, "FUNCTION APPLICATION", "mathematical invisible operator is invisible and forces UCS-2", true
	case 0x2062:
		return InvisibleMathInvisible, "INVISIBLE TIMES", "mathematical invisible operator is invisible and forces UCS-2", true
	case 0x2063:
		return InvisibleMathInvisible, "INVISIBLE SEPARATOR", "mathematical invisible separator is invisible and forces UCS-2", true
	case 0x2064:
		return InvisibleMathInvisible, "INVISIBLE PLUS", "mathematical invisible operator is invisible and forces UCS-2", true
	case 0x2065:
		return InvisibleFormatControl, "INHIBIT SYMMETRIC SWAPPING", "reserved/deprecated invisible format character", true
	case 0x2066:
		return InvisibleBidiControl, "LEFT-TO-RIGHT ISOLATE", "bidirectional isolate is invisible and can change display order", true
	case 0x2067:
		return InvisibleBidiControl, "RIGHT-TO-LEFT ISOLATE", "bidirectional isolate is invisible and can change display order", true
	case 0x2068:
		return InvisibleBidiControl, "FIRST STRONG ISOLATE", "bidirectional isolate is invisible and can change display order", true
	case 0x2069:
		return InvisibleBidiControl, "POP DIRECTIONAL ISOLATE", "bidirectional isolate terminator is invisible", true
	case 0x206A:
		return InvisibleFormatControl, "INHIBIT SYMMETRIC SWAPPING", "deprecated invisible format character", true
	case 0x206B:
		return InvisibleFormatControl, "ACTIVATE SYMMETRIC SWAPPING", "deprecated invisible format character", true
	case 0x206C:
		return InvisibleFormatControl, "INHIBIT ARABIC FORM SHAPING", "deprecated invisible format character", true
	case 0x206D:
		return InvisibleFormatControl, "ACTIVATE ARABIC FORM SHAPING", "deprecated invisible format character", true
	case 0x206E:
		return InvisibleFormatControl, "NATIONAL DIGIT SHAPES", "deprecated invisible format character", true
	case 0x206F:
		return InvisibleFormatControl, "NOMINAL DIGIT SHAPES", "deprecated invisible format character", true
	case 0x3000:
		return InvisibleUnicodeSpace, "IDEOGRAPHIC SPACE", "wide Unicode space looks blank and forces UCS-2", true
	case 0xFEFF:
		return InvisibleZeroWidth, "ZERO WIDTH NO-BREAK SPACE / BYTE ORDER MARK", "BOM/zero-width no-break space is invisible and forces UCS-2", true
	}
	if r >= 0xFE00 && r <= 0xFE0F {
		return InvisibleVariationSelector, fmt.Sprintf("VARIATION SELECTOR-%d", int(r-0xFE00+1)), "variation selector is invisible and changes emoji/text rendering", true
	}
	if r >= 0xE0100 && r <= 0xE01EF {
		return InvisibleVariationSelector, fmt.Sprintf("VARIATION SELECTOR-%d", int(r-0xE0100+17)), "supplementary variation selector is invisible and changes rendering", true
	}
	return InvisibleNone, "", "", false
}

func IsInvisibleRune(r rune) bool { _, _, _, ok := InvisibleInfo(r); return ok }
func IsZeroWidth(r rune) bool {
	c, _, _, ok := InvisibleInfo(r)
	return ok && c == InvisibleZeroWidth
}
func IsBidiControl(r rune) bool {
	c, _, _, ok := InvisibleInfo(r)
	return ok && c == InvisibleBidiControl
}
func IsFormatControl(r rune) bool {
	c, _, _, ok := InvisibleInfo(r)
	return ok && c == InvisibleFormatControl
}
func IsUnicodeSpaceSeparator(r rune) bool {
	c, _, _, ok := InvisibleInfo(r)
	return ok && c == InvisibleUnicodeSpace
}
func IsMathInvisible(r rune) bool {
	c, _, _, ok := InvisibleInfo(r)
	return ok && c == InvisibleMathInvisible
}

func AppendInvisibleCharacterPositions(dst []InvisibleCharacterPosition, text string) []InvisibleCharacterPosition {
	ri, u16 := 0, 0
	for bi, r := range text {
		if class, name, reason, ok := InvisibleInfo(r); ok {
			dst = append(dst, InvisibleCharacterPosition{IndexByte: bi, IndexRune: ri, IndexUTF16: u16, Char: string(r), Visible: VisibleChar(r), CodePoint: CodePoint(r), Hex: fmt.Sprintf("%04X", r), Name: name, Class: class, ClassLabel: class.String(), Reason: reason})
		}
		ri++
		u16 += UCS2Units(r)
	}
	return dst
}

func InvisibleCharacterPositions(text string) []InvisibleCharacterPosition {
	return AppendInvisibleCharacterPositions(nil, text)
}

func FirstInvisibleCharacterPosition(text string) *InvisibleCharacterPosition {
	ri, u16 := 0, 0
	for bi, r := range text {
		if class, name, reason, ok := InvisibleInfo(r); ok {
			c := InvisibleCharacterPosition{IndexByte: bi, IndexRune: ri, IndexUTF16: u16, Char: string(r), Visible: VisibleChar(r), CodePoint: CodePoint(r), Hex: fmt.Sprintf("%04X", r), Name: name, Class: class, ClassLabel: class.String(), Reason: reason}
			return &c
		}
		ri++
		u16 += UCS2Units(r)
	}
	return nil
}
