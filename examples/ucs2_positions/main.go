package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "Hello नमस्ते 🙂"
	for _, c := range smslite.UCS2CharacterPositions(text) {
		fmt.Printf("rune=%d byte=%d utf16=%d char=%q visible=%s code=%s reason=%s\n", c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint, c.Reason)
	}
}
