package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "A\u205F\u2060\u2061\u2062\u2063\u2064\u2065\u2066\u2067\u2068\u2069\u206A\u206B\u206C\u206D\u206E\u206FB"
	for _, c := range smslite.InvisibleCharacterPositions(text) {
		fmt.Printf("rune=%d byte=%d utf16=%d char=%q visible=%s code=%s class=%s name=%s\n", c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint, c.ClassLabel, c.Name)
	}
}
