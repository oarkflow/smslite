package smslite

import "unicode/utf8"

// asciiGSM7Cost stores GSM-7 septet cost for ASCII bytes.
// 0 means not representable in GSM-7, 1 means default table, 2 means extension table.
var asciiGSM7Cost = [128]uint8{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 1,
	0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 0,
}

func gsm7ASCIICost(b byte) uint8 {
	if b < 128 {
		return asciiGSM7Cost[b]
	}
	return 0
}

func decodeRuneAt(s string, i int) (rune, int) {
	if s[i] < utf8.RuneSelf {
		return rune(s[i]), 1
	}
	r, sz := utf8.DecodeRuneInString(s[i:])
	return r, sz
}
