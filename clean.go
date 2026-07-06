package smslite

import (
	"strings"
	"unicode"
)

// GSMCleanResult is the strict cleaning result for pasted SMS text. Cleaned is
// guaranteed to contain only GSM-7 default/extension characters: known Unicode
// punctuation, spaces, accents, mathematical letters, fullwidth forms,
// Greek/Cyrillic/Latin homoglyphs, and zero-width controls are mapped or removed;
// unknown non-GSM characters are dropped.
type GSMCleanResult struct {
	Original     string       `json:"original"`
	Cleaned      string       `json:"cleaned"`
	Changed      bool         `json:"changed"`
	Encoding     Encoding     `json:"encoding"`
	Mapped       int          `json:"mapped"`
	Removed      int          `json:"removed"`
	Replacements []Suggestion `json:"replacements,omitempty"`
}

// CleanTextToGSM returns text that is safe for GSM-7 SMS encoding. Characters
// with a known GSM-safe replacement are mapped; zero-width/invisible/default-
// ignorable characters are removed; any remaining non-GSM character is removed.
func CleanTextToGSM(text string) string {
	if text == "" {
		return ""
	}
	var b strings.Builder
	changed := false
	for i, r := range text {
		repl, keep, _ := cleanRuneForGSM7(r)
		if !keep {
			repl = ""
		}
		if repl == string(r) {
			if changed {
				b.WriteRune(r)
			}
			continue
		}
		if !changed {
			b.Grow(len(text))
			b.WriteString(text[:i])
			changed = true
		}
		b.WriteString(repl)
	}
	if !changed {
		return text
	}
	return b.String()
}

// CleanGSM is a short alias for CleanTextToGSM.
func CleanGSM(text string) string { return CleanTextToGSM(text) }

// CleanToGSM converts pasted UCS-2-triggering text into GSM-safe text.
// It is an alias for CleanTextToGSM kept for package users who prefer the
// shorter API name. Unknown unsupported characters are removed.
func CleanToGSM(text string) string { return CleanTextToGSM(text) }

// RequiresUCS2 reports whether the original text contains any character that
// would force UCS-2 before cleaning.
func RequiresUCS2(text string) bool { return !IsGSM7(text) }

// CleanForGSM7 returns the cleaned text and a detailed replacement/removal log.
func CleanForGSM7(text string) GSMCleanResult {
	if text == "" {
		return GSMCleanResult{Original: text, Cleaned: text, Encoding: EncodingGSM7}
	}
	var b strings.Builder
	var sugs []Suggestion
	changed := false
	mapped, removed := 0, 0
	ri := 0
	for i, r := range text {
		repl, keep, reason := cleanRuneForGSM7(r)
		if !keep {
			repl = ""
			reason = "drop character because no GSM-7-safe mapping is known"
		}
		if repl == string(r) {
			if changed {
				b.WriteRune(r)
			}
			ri++
			continue
		}
		if !changed {
			b.Grow(len(text))
			b.WriteString(text[:i])
			changed = true
		}
		b.WriteString(repl)
		saves := UCS2UnitLength(string(r))
		if repl != "" {
			if c, err := GSM7SeptetLength(repl); err == nil {
				saves -= c
			}
		}
		if repl == "" {
			removed++
		} else {
			mapped++
		}
		sugs = append(sugs, Suggestion{IndexRune: ri, Original: string(r), Replacement: repl, Reason: reason, SavesUnits: saves})
		ri++
	}
	cleaned := text
	if changed {
		cleaned = b.String()
	}
	enc := EncodingGSM7
	if !IsGSM7(cleaned) {
		enc = EncodingUCS2
	}
	return GSMCleanResult{Original: text, Cleaned: cleaned, Changed: changed, Encoding: enc, Mapped: mapped, Removed: removed, Replacements: sugs}
}

// AppendCleanForGSM7 appends the cleaned GSM-7 text to dst and returns the
// extended slice. It is useful on hot paths where the caller controls reuse.
func AppendCleanForGSM7(dst []byte, text string) []byte {
	for _, r := range text {
		repl, keep, _ := cleanRuneForGSM7(r)
		if !keep || repl == "" {
			continue
		}
		dst = append(dst, repl...)
	}
	return dst
}

// GSM7CleanReplacement returns the exact replacement that CleanTextToGSM would
// use for one rune. ok=false means the rune has no known mapping and is dropped.
func GSM7CleanReplacement(r rune) (replacement string, ok bool) {
	repl, keep, _ := cleanRuneForGSM7(r)
	if keep {
		return repl, true
	}
	return "", false
}

func cleanRuneForGSM7(r rune) (replacement string, keep bool, reason string) {
	if repl, ok := manualGSMCleanReplacementFor(r); ok {
		return repl, true, "map Unicode punctuation, symbol, or homoglyph to GSM-7-safe text"
	}
	if isGSMCleanDropRune(r) {
		return "", true, "remove zero-width, invisible, control, variation, or combining character"
	}
	if repl, ok := unicodeCompatibilityReplacementFor(r); ok {
		return repl, true, "map Unicode compatibility character to GSM-7-safe text"
	}
	if repl, ok := unicodeExtraCleanReplacementFor(r); ok {
		return repl, true, "map Unicode decimal digit, regional indicator, or Latin variant to GSM-7-safe text"
	}
	if _, ok := GSM7CharCost(r); ok {
		return string(r), true, ""
	}
	return "", false, ""
}

func isGSMCleanDropRune(r rune) bool {
	if r == utf8RuneError {
		return true
	}
	if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r) {
		return true
	}
	if unicode.Is(unicode.Cf, r) {
		return true
	}
	if r < 0x20 && r != '\n' && r != '\r' && r != '\f' {
		return true
	}
	if r == 0x7F {
		return true
	}
	if IsZeroWidth(r) || IsBidiControl(r) || IsFormatControl(r) || IsMathInvisible(r) {
		return true
	}
	if r >= 0xFE00 && r <= 0xFE0F || r >= 0xE0100 && r <= 0xE01EF {
		return true
	}
	switch r {
	case 0x00AD, // SOFT HYPHEN
		0x034F, // COMBINING GRAPHEME JOINER
		0x061C, // ARABIC LETTER MARK
		0x115F, // HANGUL CHOSEONG FILLER
		0x1160, // HANGUL JUNGSEONG FILLER
		0x17B4, // KHMER VOWEL INHERENT AQ
		0x17B5, // KHMER VOWEL INHERENT AA
		0x180B, // MONGOLIAN FREE VARIATION SELECTOR ONE
		0x180C, // MONGOLIAN FREE VARIATION SELECTOR TWO
		0x180D, // MONGOLIAN FREE VARIATION SELECTOR THREE
		0x180E, // MONGOLIAN VOWEL SEPARATOR
		0x180F, // MONGOLIAN FREE VARIATION SELECTOR FOUR
		0x200B, // ZERO WIDTH SPACE
		0x200C, // ZERO WIDTH NON-JOINER
		0x200D, // ZERO WIDTH JOINER
		0x200E, // LEFT-TO-RIGHT MARK
		0x200F, // RIGHT-TO-LEFT MARK
		0x202A, // LEFT-TO-RIGHT EMBEDDING
		0x202B, // RIGHT-TO-LEFT EMBEDDING
		0x202C, // POP DIRECTIONAL FORMATTING
		0x202D, // LEFT-TO-RIGHT OVERRIDE
		0x202E, // RIGHT-TO-LEFT OVERRIDE
		0x2060, // WORD JOINER
		0x2061, // FUNCTION APPLICATION
		0x2062, // INVISIBLE TIMES
		0x2063, // INVISIBLE SEPARATOR
		0x2064, // INVISIBLE PLUS
		0x2065, // INHIBIT SYMMETRIC SWAPPING
		0x2066, // LEFT-TO-RIGHT ISOLATE
		0x2067, // RIGHT-TO-LEFT ISOLATE
		0x2068, // FIRST STRONG ISOLATE
		0x2069, // POP DIRECTIONAL ISOLATE
		0x206A, // INHIBIT SYMMETRIC SWAPPING (deprecated)
		0x206B, // ACTIVATE SYMMETRIC SWAPPING (deprecated)
		0x206C, // INHIBIT ARABIC FORM SHAPING (deprecated)
		0x206D, // ACTIVATE ARABIC FORM SHAPING (deprecated)
		0x206E, // NATIONAL DIGIT SHAPES (deprecated)
		0x206F, // NOMINAL DIGIT SHAPES (deprecated)
		0x3164, // HANGUL FILLER
		0xFFA0, // HALFWIDTH HANGUL FILLER
		0xFEFF, // ZERO WIDTH NO-BREAK SPACE / BOM
		0xFFF9, // INTERLINEAR ANNOTATION ANCHOR
		0xFFFA, // INTERLINEAR ANNOTATION SEPARATOR
		0xFFFB: // INTERLINEAR ANNOTATION TERMINATOR
		return true
	}
	if r >= 0x1BCA0 && r <= 0x1BCA3 { // SHORTHAND FORMAT CONTROLS
		return true
	}
	if r >= 0x1D173 && r <= 0x1D17A { // MUSICAL SYMBOL FORMAT CONTROLS
		return true
	}
	if r >= 0xE0000 && r <= 0xE007F { // LANGUAGE/TAG characters, including CANCEL TAG
		return true
	}
	return false
}

const utf8RuneError = '\uFFFD'

func manualGSMCleanReplacementFor(r rune) (string, bool) {
	if unicode.Is(unicode.Zs, r) {
		return " ", true
	}
	switch r {
	// ASCII-ish whitespace and separators.
	case '\t', '\v':
		return " ", true
	case 0x00A0, 0x1680, 0x2000, 0x2001, 0x2002, 0x2003, 0x2004, 0x2005, 0x2006, 0x2007, 0x2008, 0x2009, 0x200A, 0x202F, 0x205F, 0x3000:
		return " ", true
	case 0x2028, 0x2029:
		return "\n", true

	// Quotes, apostrophes, and accents commonly pasted from rich text.
	case '`', '´', 'ʻ', 'ʼ', 'ʹ', 'ʽ', 'ʾ', 'ʿ', 'ˈ', 'ˊ', 'ˋ', '˴', '˵', '՚', 'ꞌ', '＇', '‘', '’', '‚', '‛', '′', '‵', '❛', '❜':
		return "'", true
	case '“', '”', '„', '‟', '″', '‶', '❝', '❞', '〝', '〞', '＂':
		return "\"", true
	case '‹', '❮', '〈', '《', '«':
		return "<", true
	case '›', '❯', '〉', '》', '»':
		return ">", true

	// Dashes, hyphens, minus signs, and line-drawing variants.
	case '‐', '‑', '‒', '–', '—', '―', '−', '﹘', '﹣', '－', '⁃':
		return "-", true
	case '…':
		return "...", true
	case '•', '·', '∙', '●', '○', '◦', '▪', '▫', '■', '□', '◆', '◇', '★', '☆':
		return "*", true
	case '×', '✕', '✖', '✗', '✘':
		return "x", true
	case '÷':
		return "/", true
	case '∕', '⁄', '╱', '／':
		return "/", true
	case '＼', '╲':
		return "\\", true
	case '＿':
		return "_", true

	// Fractions and common number symbols.
	case '¼':
		return "1/4", true
	case '½':
		return "1/2", true
	case '¾':
		return "3/4", true
	case '⅐':
		return "1/7", true
	case '⅑':
		return "1/9", true
	case '⅒':
		return "1/10", true
	case '⅓':
		return "1/3", true
	case '⅔':
		return "2/3", true
	case '⅕':
		return "1/5", true
	case '⅖':
		return "2/5", true
	case '⅗':
		return "3/5", true
	case '⅘':
		return "4/5", true
	case '⅙':
		return "1/6", true
	case '⅚':
		return "5/6", true
	case '⅛':
		return "1/8", true
	case '⅜':
		return "3/8", true
	case '⅝':
		return "5/8", true
	case '⅞':
		return "7/8", true

	// Currency and business symbols not in GSM-7.
	case '₹', '₨':
		return "Rs", true
	case '₽':
		return "R", true
	case '₩':
		return "W", true
	case '₦':
		return "N", true
	case '₱':
		return "P", true
	case '₫':
		return "d", true
	case '₭':
		return "K", true
	case '₮':
		return "T", true
	case '₴':
		return "UAH", true
	case '₿':
		return "BTC", true
	case '¢':
		return "c", true
	case '©':
		return "(c)", true
	case '®':
		return "(r)", true
	case '™':
		return "TM", true
	case '℠':
		return "SM", true
	case '°':
		return " deg ", true
	case '№':
		return "No", true

	// Latin letters that do not decompose cleanly with NFKD.
	case 'Æ', 'Ǣ', 'Ǽ', 'ᴁ':
		return "AE", true
	case 'æ', 'ǣ', 'ǽ':
		return "ae", true
	case 'Œ', 'ɶ', 'Ꝏ':
		return "OE", true
	case 'œ', 'ꝏ':
		return "oe", true
	case 'ß', 'ẞ':
		return "ss", true
	case 'Ð', 'Đ', 'Ɖ', 'Ɗ', 'ᴅ', 'ꓓ':
		return "D", true
	case 'ð', 'đ', 'ɗ', 'ɖ':
		return "d", true
	case 'Þ', 'Ƥ':
		return "P", true
	case 'þ', 'ƥ':
		return "p", true
	case 'Ł', 'Ŀ', 'Ƚ', 'Ⱡ', 'Ꝉ':
		return "L", true
	case 'ł', 'ŀ', 'ƚ', 'ꝉ':
		return "l", true
	case 'Ŋ', 'Ɲ', 'Ꞑ':
		return "N", true
	case 'ŋ', 'ɲ', 'ꞑ':
		return "n", true
	case 'Ŧ', 'Ƭ', 'Ʈ', 'Ⱦ', 'Ꞇ':
		return "T", true
	case 'ŧ', 'ƭ', 'ʈ', 'ⱦ', 'ꞇ':
		return "t", true
	case 'Ɓ', 'Ƀ', 'Ꞗ', 'ꓐ':
		return "B", true
	case 'ɓ', 'ƀ', 'ƃ':
		return "b", true
	case 'Ƈ', 'Ȼ', 'Ꞓ', 'ꓚ':
		return "C", true
	case 'ƈ', 'ȼ', 'ꞓ':
		return "c", true
	case 'Ƒ', 'Ꞙ':
		return "F", true
	case 'ƒ', 'ꞙ':
		return "f", true
	case 'Ɠ', 'Ǥ', 'Ꞡ':
		return "G", true
	case 'ɠ', 'ǥ', 'ᵹ', 'ꞡ':
		return "g", true
	case 'Ƕ', 'Ɦ':
		return "H", true
	case 'ƕ', 'ꞕ', 'Ꞃ', 'ꞃ':
		return "h", true
	case 'Ɨ', 'Ɩ', 'Ǐ', 'Ɪ', 'ꓲ':
		return "I", true
	case 'ɨ', 'ɩ', 'ǐ', 'ı', 'ꭵ':
		return "i", true
	case 'Ɉ', 'Ʝ', 'ꓙ':
		return "J", true
	case 'ɉ', 'ȷ', 'ɟ':
		return "j", true
	case 'Ƙ', 'Ꝁ', 'ꓗ':
		return "K", true
	case 'ƙ', 'ꝁ':
		return "k", true
	case 'Ɱ', 'Ɯ', 'ꟽ', 'ꓟ':
		return "M", true
	case 'ɯ', 'ɰ':
		return "m", true
	case 'Ɵ', 'Ø', 'Ǿ', 'Ꝋ', 'ꓳ':
		return "O", true
	case 'ɵ', 'ø', 'ǿ', 'ꝋ':
		return "o", true
	case 'Ɍ', 'Ꞧ', 'ꓣ':
		return "R", true
	case 'ɍ', 'ɽ', 'ꞧ':
		return "r", true
	case 'Ƨ', 'Ꞩ', 'ꓢ':
		return "S", true
	case 'ƨ', 'ʂ', 'ꞩ':
		return "s", true
	case 'Ʊ', 'Ư', 'ꓵ':
		return "U", true
	case 'ʊ', 'ư':
		return "u", true
	case 'Ʋ', 'Ꝟ', 'ꓦ':
		return "V", true
	case 'ʋ', 'ꝟ':
		return "v", true
	case 'Ƿ', 'Ꟃ', 'ꓪ':
		return "W", true
	case 'ƿ', 'ʍ':
		return "w", true
	case 'Ƴ', 'Ȳ', 'Ɏ', 'ꓬ':
		return "Y", true
	case 'ƴ', 'ȳ', 'ɏ':
		return "y", true
	case 'Ƶ', 'Ȥ', 'Ⱬ', 'ꓜ':
		return "Z", true
	case 'ƶ', 'ȥ', 'ɀ', 'ʐ', 'ʑ':
		return "z", true
	}
	if repl, ok := greekCleanReplacementFor(r); ok {
		return repl, true
	}
	if repl, ok := cyrillicCleanReplacementFor(r); ok {
		return repl, true
	}
	if repl, ok := letterlikeCleanReplacementFor(r); ok {
		return repl, true
	}
	return "", false
}

func greekCleanReplacementFor(r rune) (string, bool) {
	switch r {
	case 'Α', 'Ά', 'Ἀ', 'Ἁ', 'Ἂ', 'Ἃ', 'Ἄ', 'Ἅ', 'Ἆ', 'Ἇ', 'ᾈ', 'ᾉ', 'ᾊ', 'ᾋ', 'ᾌ', 'ᾍ', 'ᾎ', 'ᾏ', 'Ᾰ', 'Ᾱ', 'Ὰ', 'Ά', 'ᾼ':
		return "A", true
	case 'α', 'ά', 'ἀ', 'ἁ', 'ἂ', 'ἃ', 'ἄ', 'ἅ', 'ἆ', 'ἇ', 'ᾀ', 'ᾁ', 'ᾂ', 'ᾃ', 'ᾄ', 'ᾅ', 'ᾆ', 'ᾇ', 'ὰ', 'ά', 'ᾰ', 'ᾱ', 'ᾲ', 'ᾳ', 'ᾴ', 'ᾶ', 'ᾷ':
		return "a", true
	case 'Β':
		return "B", true
	case 'β', 'ϐ':
		return "b", true
	case 'Γ':
		return "G", true
	case 'γ':
		return "g", true
	case 'Δ':
		return "D", true
	case 'δ':
		return "d", true
	case 'Ε', 'Έ', 'Ἐ', 'Ἑ', 'Ἒ', 'Ἓ', 'Ἔ', 'Ἕ', 'Ὲ', 'Έ':
		return "E", true
	case 'ε', 'έ', 'ϵ', '϶', 'ἐ', 'ἑ', 'ἒ', 'ἓ', 'ἔ', 'ἕ', 'ὲ', 'έ':
		return "e", true
	case 'Ζ':
		return "Z", true
	case 'ζ':
		return "z", true
	case 'Η', 'Ή', 'Ἠ', 'Ἡ', 'Ἢ', 'Ἣ', 'Ἤ', 'Ἥ', 'Ἦ', 'Ἧ', 'ᾘ', 'ᾙ', 'ᾚ', 'ᾛ', 'ᾜ', 'ᾝ', 'ᾞ', 'ᾟ', 'Ὴ', 'Ή', 'ῌ':
		return "H", true
	case 'η', 'ή', 'ἠ', 'ἡ', 'ἢ', 'ἣ', 'ἤ', 'ἥ', 'ἦ', 'ἧ', 'ᾐ', 'ᾑ', 'ᾒ', 'ᾓ', 'ᾔ', 'ᾕ', 'ᾖ', 'ᾗ', 'ὴ', 'ή', 'ῂ', 'ῃ', 'ῄ', 'ῆ', 'ῇ':
		return "n", true
	case 'Θ', 'ϴ':
		return "O", true
	case 'θ', 'ϑ':
		return "o", true
	case 'Ι', 'Ί', 'Ϊ', 'Ἰ', 'Ἱ', 'Ἲ', 'Ἳ', 'Ἴ', 'Ἵ', 'Ἶ', 'Ἷ', 'Ῐ', 'Ῑ', 'Ὶ', 'Ί':
		return "I", true
	case 'ι', 'ί', 'ϊ', 'ΐ', 'ἰ', 'ἱ', 'ἲ', 'ἳ', 'ἴ', 'ἵ', 'ἶ', 'ἷ', 'ὶ', 'ί', 'ῐ', 'ῑ', 'ῒ', 'ΐ', 'ῖ', 'ῗ', 'ͺ':
		return "i", true
	case 'Κ', 'Ϗ':
		return "K", true
	case 'κ', 'ϰ':
		return "k", true
	case 'Λ':
		return "L", true
	case 'λ':
		return "l", true
	case 'Μ':
		return "M", true
	case 'μ', 'µ':
		return "u", true
	case 'Ν':
		return "N", true
	case 'ν':
		return "v", true
	case 'Ξ':
		return "X", true
	case 'ξ':
		return "x", true
	case 'Ο', 'Ό', 'Ὀ', 'Ὁ', 'Ὂ', 'Ὃ', 'Ὄ', 'Ὅ', 'Ὸ', 'Ό':
		return "O", true
	case 'ο', 'ό', 'σ', 'ὀ', 'ὁ', 'ὂ', 'ὃ', 'ὄ', 'ὅ', 'ὸ', 'ό':
		return "o", true
	case 'Π', 'ϖ':
		return "P", true
	case 'π':
		return "p", true
	case 'Ρ', 'Ῥ':
		return "P", true
	case 'ρ', 'ῤ', 'ῥ':
		return "p", true
	case 'Σ', 'Ϲ':
		return "S", true
	case 'ς', 'ϲ':
		return "c", true
	case 'Τ':
		return "T", true
	case 'τ':
		return "t", true
	case 'Υ', 'Ύ', 'Ϋ', 'ϒ', 'Ὑ', 'Ὓ', 'Ὕ', 'Ὗ', 'Ῠ', 'Ῡ', 'Ὺ', 'Ύ':
		return "Y", true
	case 'υ', 'ύ', 'ϋ', 'ΰ', 'ὐ', 'ὑ', 'ὒ', 'ὓ', 'ὔ', 'ὕ', 'ὖ', 'ὗ', 'ὺ', 'ύ', 'ῠ', 'ῡ', 'ῢ', 'ΰ', 'ῦ', 'ῧ':
		return "u", true
	case 'Φ':
		return "F", true
	case 'φ', 'ϕ':
		return "f", true
	case 'Χ':
		return "X", true
	case 'χ':
		return "x", true
	case 'Ψ':
		return "Y", true
	case 'ψ':
		return "y", true
	case 'Ω', 'Ώ', 'Ω', 'Ὠ', 'Ὡ', 'Ὢ', 'Ὣ', 'Ὤ', 'Ὥ', 'Ὦ', 'Ὧ', 'ᾨ', 'ᾩ', 'ᾪ', 'ᾫ', 'ᾬ', 'ᾭ', 'ᾮ', 'ᾯ', 'Ὼ', 'Ώ', 'ῼ':
		return "O", true
	case 'ω', 'ώ', 'ὠ', 'ὡ', 'ὢ', 'ὣ', 'ὤ', 'ὥ', 'ὦ', 'ὧ', 'ᾠ', 'ᾡ', 'ᾢ', 'ᾣ', 'ᾤ', 'ᾥ', 'ᾦ', 'ᾧ', 'ὼ', 'ώ', 'ῲ', 'ῳ', 'ῴ', 'ῶ', 'ῷ':
		return "w", true
	}
	return "", false
}

func cyrillicCleanReplacementFor(r rune) (string, bool) {
	switch r {
	case 'А', 'Ӑ', 'Ӓ':
		return "A", true
	case 'а', 'ӑ', 'ӓ':
		return "a", true
	case 'В', 'Ᏼ':
		return "B", true
	case 'в':
		return "b", true
	case 'С', 'Ϲ', 'Ҫ':
		return "C", true
	case 'с', 'ϲ', 'ҫ':
		return "c", true
	case 'Е', 'Ё', 'Є', 'Ӗ':
		return "E", true
	case 'е', 'ё', 'є', 'ӗ':
		return "e", true
	case 'Н', 'Ң', 'Ҥ':
		return "H", true
	case 'н', 'ң', 'ҥ':
		return "h", true
	case 'І', 'Ї', 'Ӏ':
		return "I", true
	case 'і', 'ї', 'ӏ':
		return "i", true
	case 'Ј':
		return "J", true
	case 'ј':
		return "j", true
	case 'К', 'Ҡ', 'Қ', 'Ҝ':
		return "K", true
	case 'к', 'ҡ', 'қ', 'ҝ':
		return "k", true
	case 'М':
		return "M", true
	case 'м':
		return "m", true
	case 'О', 'Ӧ', 'Ө', 'Ӫ':
		return "O", true
	case 'о', 'ӧ', 'ө', 'ӫ':
		return "o", true
	case 'Р':
		return "P", true
	case 'р':
		return "p", true
	case 'Т', 'Ҭ':
		return "T", true
	case 'т', 'ҭ':
		return "t", true
	case 'У', 'Ү', 'Ұ', 'Ӯ', 'Ӱ', 'Ӳ':
		return "Y", true
	case 'у', 'ү', 'ұ', 'ӯ', 'ӱ', 'ӳ':
		return "y", true
	case 'Х', 'Ҳ':
		return "X", true
	case 'х', 'ҳ':
		return "x", true
	case 'Ѕ':
		return "S", true
	case 'ѕ':
		return "s", true
	case 'Ԍ':
		return "G", true
	case 'ԍ':
		return "g", true
	case 'Ԛ':
		return "Q", true
	case 'ԛ':
		return "q", true
	case 'Ԝ':
		return "W", true
	case 'ԝ':
		return "w", true
	case 'Ӕ':
		return "AE", true
	case 'ӕ':
		return "ae", true
	case 'Б', 'Ь':
		return "B", true
	case 'б', 'ь':
		return "b", true
	case 'Г':
		return "G", true
	case 'г':
		return "r", true
	case 'Д':
		return "D", true
	case 'д':
		return "d", true
	case 'З':
		return "Z", true
	case 'з':
		return "z", true
	case 'Л':
		return "L", true
	case 'л':
		return "l", true
	case 'П':
		return "P", true
	case 'п':
		return "n", true
	case 'Ф':
		return "O", true
	case 'ф':
		return "o", true
	case 'Ч':
		return "Y", true
	case 'ч':
		return "y", true
	case 'Ш', 'Щ':
		return "W", true
	case 'ш', 'щ':
		return "w", true
	case 'Э':
		return "E", true
	case 'э':
		return "e", true
	case 'Ю':
		return "IO", true
	case 'ю':
		return "io", true
	case 'Я':
		return "R", true
	case 'я':
		return "r", true
	}
	return "", false
}

func letterlikeCleanReplacementFor(r rune) (string, bool) {
	switch r {
	case 'ℂ', 'ℭ', '℃':
		return "C", true
	case '℅':
		return "c/o", true
	case '℆':
		return "c/u", true
	case 'ℊ':
		return "g", true
	case 'ℋ', 'ℌ', 'ℍ':
		return "H", true
	case 'ℎ':
		return "h", true
	case 'ℐ', 'ℑ', 'ℓ', 'ℒ':
		return "I", true
	case 'ℕ':
		return "N", true
	case 'ℙ':
		return "P", true
	case 'ℚ':
		return "Q", true
	case 'ℛ', 'ℜ', 'ℝ':
		return "R", true
	case '℡':
		return "TEL", true
	case 'ℤ':
		return "Z", true
	case 'K':
		return "K", true
	case 'Å':
		return "A", true
	case 'ℬ':
		return "B", true
	case 'ℰ':
		return "E", true
	case 'ℱ':
		return "F", true
	case 'ℳ':
		return "M", true
	case 'ℴ':
		return "o", true
	case 'ℹ':
		return "i", true
	}
	return "", false
}
