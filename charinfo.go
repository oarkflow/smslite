package smslite

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func BuildCharacterInfo(text string) []CharacterInfo {
	out := make([]CharacterInfo, 0, utf8.RuneCountInString(text))
	return AppendCharacterInfo(out, text)
}

func AppendCharacterInfo(dst []CharacterInfo, text string) []CharacterInfo {
	ri, u16 := 0, 0
	for bi, r := range text {
		ci := CharacterInfoForRune(r, bi, ri, u16)
		dst = append(dst, ci)
		ri++
		u16 += UCS2Units(r)
	}
	return dst
}

func CharacterInfoForRune(r rune, byteIndex, runeIndex, utf16Index int) CharacterInfo {
	cost, gsm := GSM7CharCost(r)
	isDef := IsGSM7Default(r)
	isExt := IsGSM7Extended(r)
	emoji := IsEmoji(r)
	combining := unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r)
	variation := r >= 0xFE00 && r <= 0xFE0F || r >= 0xE0100 && r <= 0xE01EF
	zwj := r == 0x200D
	invClass, invName, invReason, invis := InvisibleInfo(r)
	kind := KindUCS2
	enc := EncodingUCS2
	if isDef {
		kind, enc = KindGSMDefault, EncodingGSM7
	}
	if isExt {
		kind, enc = KindGSMExtended, EncodingGSM7
	}
	if emoji {
		kind = KindEmoji
	}
	if combining {
		kind = KindCombining
	}
	if invis {
		kind = KindInvisible
	}
	if unicode.IsControl(r) && !isDef {
		kind = KindControl
	}
	if !gsm {
		cost = 0
	}
	cp := CodePoint(r)
	visible := VisibleChar(r)
	name := UnicodeName(r)
	if name == "" && invName != "" {
		name = invName
	}
	warning := ""
	if invReason != "" {
		warning = invReason
	}
	return CharacterInfo{
		IndexByte: byteIndex, IndexRune: runeIndex, IndexUTF16: utf16Index,
		Char: string(r), Rune: r, CodePoint: cp, UnicodeName: name, Category: UnicodeCategory(r),
		Kind: kind, KindLabel: kind.String(), Encoding: enc, IsGSMDefault: isDef, IsGSMExtended: isExt,
		IsUCS2: !gsm, IsEmoji: emoji, IsCombiningMark: combining, IsVariation: variation, IsZeroWidthJoiner: zwj,
		IsZeroWidth: IsZeroWidth(r), IsBidiControl: IsBidiControl(r), IsFormatControl: IsFormatControl(r), IsMathInvisible: IsMathInvisible(r),
		InvisibleClass: invClass, InvisibleClassLabel: invClass.String(), IsInvisible: invis, GSMSeptets: cost, UCS2Units: UCS2Units(r), Visible: visible, DisplayLabel: visible, Hex: fmt.Sprintf("%04X", r), Warning: warning,
	}
}

func CodePoint(r rune) string {
	if r <= 0xFFFF {
		return fmt.Sprintf("U+%04X", r)
	}
	return fmt.Sprintf("U+%X", r)
}

func UnicodeCategory(r rune) string {
	switch {
	case unicode.IsLetter(r):
		return "Letter"
	case unicode.IsDigit(r):
		return "Number"
	case unicode.IsSpace(r):
		return "Space"
	case unicode.IsPunct(r):
		return "Punctuation"
	case unicode.IsSymbol(r):
		return "Symbol"
	case unicode.IsControl(r):
		return "Control"
	case unicode.IsMark(r):
		return "Mark"
	default:
		return "Other"
	}
}

func UnicodeName(r rune) string {
	switch r {
	case ' ':
		return "SPACE"
	case '\n':
		return "LINE FEED"
	case '\r':
		return "CARRIAGE RETURN"
	case '\t':
		return "CHARACTER TABULATION"
	case 0x00A0:
		return "NO-BREAK SPACE"
	case 0x200B:
		return "ZERO WIDTH SPACE"
	case 0x200C:
		return "ZERO WIDTH NON-JOINER"
	case 0x200D:
		return "ZERO WIDTH JOINER"
	case 0xFE0E:
		return "VARIATION SELECTOR-15"
	case 0xFE0F:
		return "VARIATION SELECTOR-16"
	case '€':
		return "EURO SIGN"
	case '“':
		return "LEFT DOUBLE QUOTATION MARK"
	case '”':
		return "RIGHT DOUBLE QUOTATION MARK"
	case '‘':
		return "LEFT SINGLE QUOTATION MARK"
	case '’':
		return "RIGHT SINGLE QUOTATION MARK"
	case '–':
		return "EN DASH"
	case '—':
		return "EM DASH"
	case '…':
		return "HORIZONTAL ELLIPSIS"
	}
	if r >= 'A' && r <= 'Z' {
		return "LATIN CAPITAL LETTER " + string(r)
	}
	if r >= 'a' && r <= 'z' {
		return "LATIN SMALL LETTER " + string(r-'a'+'A')
	}
	if r >= '0' && r <= '9' {
		return "DIGIT " + string(r)
	}
	if IsEmoji(r) {
		return "EMOJI"
	}
	return ""
}

func IsInvisible(r rune) bool { return IsInvisibleRune(r) }

func VisibleChar(r rune) string {
	switch r {
	case ' ':
		return "␠"
	case '\t':
		return "⇥"
	case '\n':
		return "↵"
	case '\r':
		return "␍"
	case '\v':
		return "␋"
	case '\f':
		return "␌"
	case 0x00A0:
		return "⟦NBSP⟧"
	case 0x1680:
		return "⟦OGHAM-SP⟧"
	case 0x180E:
		return "⟦MVS⟧"
	case 0x2000:
		return "⟦EN-QUAD⟧"
	case 0x2001:
		return "⟦EM-QUAD⟧"
	case 0x2002:
		return "⟦EN-SP⟧"
	case 0x2003:
		return "⟦EM-SP⟧"
	case 0x2004:
		return "⟦3/MSP⟧"
	case 0x2005:
		return "⟦4/MSP⟧"
	case 0x2006:
		return "⟦6/MSP⟧"
	case 0x2007:
		return "⟦FIG-SP⟧"
	case 0x2008:
		return "⟦PUNC-SP⟧"
	case 0x2009:
		return "⟦THIN-SP⟧"
	case 0x200A:
		return "⟦HAIR-SP⟧"
	case 0x200B:
		return "⟦ZWSP⟧"
	case 0x200C:
		return "⟦ZWNJ⟧"
	case 0x200D:
		return "⟦ZWJ⟧"
	case 0x200E:
		return "⟦LRM⟧"
	case 0x200F:
		return "⟦RLM⟧"
	case 0x2028:
		return "⟦LS⟧"
	case 0x2029:
		return "⟦PS⟧"
	case 0x202A:
		return "⟦LRE⟧"
	case 0x202B:
		return "⟦RLE⟧"
	case 0x202C:
		return "⟦PDF⟧"
	case 0x202D:
		return "⟦LRO⟧"
	case 0x202E:
		return "⟦RLO⟧"
	case 0x202F:
		return "⟦NNBSP⟧"
	case 0x205F:
		return "⟦MMSP⟧"
	case 0x2060:
		return "⟦WJ⟧"
	case 0x2061:
		return "⟦FUNC-APP⟧"
	case 0x2062:
		return "⟦INV-TIMES⟧"
	case 0x2063:
		return "⟦INV-SEP⟧"
	case 0x2064:
		return "⟦INV-PLUS⟧"
	case 0x2065:
		return "⟦ISS⟧"
	case 0x2066:
		return "⟦LRI⟧"
	case 0x2067:
		return "⟦RLI⟧"
	case 0x2068:
		return "⟦FSI⟧"
	case 0x2069:
		return "⟦PDI⟧"
	case 0x206A:
		return "⟦ISS-OLD⟧"
	case 0x206B:
		return "⟦ASS⟧"
	case 0x206C:
		return "⟦IAFS⟧"
	case 0x206D:
		return "⟦AAFS⟧"
	case 0x206E:
		return "⟦NADS⟧"
	case 0x206F:
		return "⟦NODS⟧"
	case 0x3000:
		return "⟦IDEO-SP⟧"
	case 0xFEFF:
		return "⟦BOM/ZWNBSP⟧"
	}
	if r >= 0xFE00 && r <= 0xFE0F {
		return fmt.Sprintf("⟦VS%d⟧", int(r-0xFE00+1))
	}
	if r >= 0xE0100 && r <= 0xE01EF {
		return fmt.Sprintf("⟦VS%d⟧", int(r-0xE0100+17))
	}
	if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) {
		return "◌" + string(r)
	}
	if unicode.IsControl(r) {
		return fmt.Sprintf("⟦CTRL U+%04X⟧", r)
	}
	return string(r)
}

func IsEmoji(r rune) bool {
	return (r >= 0x1F300 && r <= 0x1FAFF) || (r >= 0x2600 && r <= 0x27BF)
}

func FindUCS2Characters(text string) []CharacterInfo {
	var out []CharacterInfo
	return AppendUCS2Characters(out, text)
}

func AppendUCS2Characters(dst []CharacterInfo, text string) []CharacterInfo {
	ri, u16 := 0, 0
	for bi, r := range text {
		if _, ok := GSM7CharCost(r); !ok {
			dst = append(dst, CharacterInfoForRune(r, bi, ri, u16))
		}
		ri++
		u16 += UCS2Units(r)
	}
	return dst
}

func HasUCS2(text string) bool { return IsUCS2Required(text) }

func FirstUCS2Character(text string) *CharacterInfo {
	ri, u16 := 0, 0
	for bi, r := range text {
		if _, ok := GSM7CharCost(r); !ok {
			c := CharacterInfoForRune(r, bi, ri, u16)
			return &c
		}
		ri++
		u16 += UCS2Units(r)
	}
	return nil
}

func UCS2Positions(text string) []int {
	var out []int
	ri := 0
	for _, r := range text {
		if _, ok := GSM7CharCost(r); !ok {
			out = append(out, ri)
		}
		ri++
	}
	return out
}

// UCS2CharacterPositions returns both the exact positions and the exact characters
// that force the entire SMS to UCS-2. Prefer this over UCS2Positions when the UI,
// API response, or log must show which character exists at each position.
func UCS2CharacterPositions(text string) []UCS2CharacterPosition {
	var out []UCS2CharacterPosition
	return AppendUCS2CharacterPositions(out, text)
}

// AppendUCS2CharacterPositions appends compact UCS-2-forcing character records to dst.
// Reusing dst lets callers keep this path allocation-free after the first capacity setup.
func AppendUCS2CharacterPositions(dst []UCS2CharacterPosition, text string) []UCS2CharacterPosition {
	ri, u16 := 0, 0
	for bi, r := range text {
		if _, ok := GSM7CharCost(r); !ok {
			dst = append(dst, UCS2CharacterPosition{
				IndexByte:  bi,
				IndexRune:  ri,
				IndexUTF16: u16,
				Char:       string(r),
				Visible:    VisibleChar(r),
				CodePoint:  CodePoint(r),
				Hex:        fmt.Sprintf("%04X", r),
				Reason:     "not present in GSM-7 default or extension alphabet; message is encoded as UCS-2",
			})
		}
		ri++
		u16 += UCS2Units(r)
	}
	return dst
}

// FirstUCS2CharacterPosition returns the first UCS-2-forcing character with indexes
// and display fields. It returns nil when the text is fully GSM-7 compatible.
func FirstUCS2CharacterPosition(text string) *UCS2CharacterPosition {
	ri, u16 := 0, 0
	for bi, r := range text {
		if _, ok := GSM7CharCost(r); !ok {
			c := UCS2CharacterPosition{IndexByte: bi, IndexRune: ri, IndexUTF16: u16, Char: string(r), Visible: VisibleChar(r), CodePoint: CodePoint(r), Hex: fmt.Sprintf("%04X", r), Reason: "not present in GSM-7 default or extension alphabet; message is encoded as UCS-2"}
			return &c
		}
		ri++
		u16 += UCS2Units(r)
	}
	return nil
}

func FindInvisibleCharacters(text string) []CharacterInfo {
	var out []CharacterInfo
	ri, u16 := 0, 0
	for bi, r := range text {
		if IsInvisible(r) {
			out = append(out, CharacterInfoForRune(r, bi, ri, u16))
		}
		ri++
		u16 += UCS2Units(r)
	}
	return out
}

func FindEmojiCharacters(text string) []CharacterInfo {
	var out []CharacterInfo
	ri, u16 := 0, 0
	for bi, r := range text {
		if IsEmoji(r) {
			out = append(out, CharacterInfoForRune(r, bi, ri, u16))
		}
		ri++
		u16 += UCS2Units(r)
	}
	return out
}

func FindNonGSMCharacters(text string) []CharacterInfo { return FindUCS2Characters(text) }
