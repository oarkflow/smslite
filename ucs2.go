package smslite

import "unicode/utf8"

func UCS2Units(r rune) int {
	if r < 0x10000 {
		return 1
	}
	return 2
}

func UCS2UnitLength(text string) int {
	n := 0
	for i := 0; i < len(text); {
		b := text[i]
		if b < 0x80 {
			n++
			i++
			continue
		}
		r, sz := decodeRuneAt(text, i)
		if r < 0x10000 {
			n++
		} else {
			n += 2
		}
		i += sz
	}
	return n
}

func RuneCount(text string) int { return utf8.RuneCountInString(text) }

func IsUCS2Required(text string) bool { return !IsGSM7(text) }
