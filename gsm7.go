package smslite

const (
	GSM7SingleLimit    = 160
	GSM7MultipartLimit = 153
	UCS2SingleLimit    = 70
	UCS2MultipartLimit = 67
	ConcatUDH8Bytes    = 6
	ConcatUDH16Bytes   = 7
	ConcatUDH8Septets  = 7
	ConcatUDH16Septets = 8
)

// IsGSM7Default reports whether r exists in GSM 03.38 default alphabet.
func IsGSM7Default(r rune) bool {
	switch r {
	case '@', '£', '$', '¥', 'è', 'é', 'ù', 'ì', 'ò', 'Ç', '\n', 'Ø', 'ø', '\r', 'Å', 'å',
		'Δ', '_', 'Φ', 'Γ', 'Λ', 'Ω', 'Π', 'Ψ', 'Σ', 'Θ', 'Ξ', 'Æ', 'æ', 'ß', 'É',
		' ', '!', '"', '#', '¤', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';', '<', '=', '>', '?',
		'¡', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'Ä', 'Ö', 'Ñ', 'Ü', '§',
		'¿', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'ä', 'ö', 'ñ', 'ü', 'à':
		return true
	default:
		return false
	}
}

// IsGSM7Extended reports whether r exists in GSM 03.38 extension table.
func IsGSM7Extended(r rune) bool {
	switch r {
	case '\f', '^', '{', '}', '\\', '[', '~', ']', '|', '€':
		return true
	default:
		return false
	}
}

func GSM7CharCost(r rune) (int, bool) {
	if r < 128 {
		c := asciiGSM7Cost[byte(r)]
		return int(c), c != 0
	}
	if IsGSM7Default(r) {
		return 1, true
	}
	if IsGSM7Extended(r) {
		return 2, true
	}
	return 0, false
}

func IsGSM7(text string) bool {
	for i := 0; i < len(text); {
		b := text[i]
		if b < 0x80 {
			if asciiGSM7Cost[b] == 0 {
				return false
			}
			i++
			continue
		}
		r, sz := decodeRuneAt(text, i)
		if _, ok := GSM7CharCost(r); !ok {
			return false
		}
		i += sz
	}
	return true
}

func GSM7SeptetLength(text string) (int, error) {
	n, ok := GSM7SeptetLengthFast(text)
	if !ok {
		return 0, ErrNonGSM7
	}
	return n, nil
}

func GSM7SeptetLengthFast(text string) (int, bool) {
	n := 0
	for i := 0; i < len(text); {
		b := text[i]
		if b < 0x80 {
			c := asciiGSM7Cost[b]
			if c == 0 {
				return n, false
			}
			n += int(c)
			i++
			continue
		}
		r, sz := decodeRuneAt(text, i)
		c, ok := GSM7CharCost(r)
		if !ok {
			return n, false
		}
		n += c
		i += sz
	}
	return n, true
}
