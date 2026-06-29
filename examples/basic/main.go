package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "Hello {name}, use code 123456 🙂"
	a := smslite.AnalyzeDetailed(text)
	fmt.Println(smslite.FormatAnalysisSummary(a))
	fmt.Println("UCS2 positions only:", smslite.UCS2Positions(text))
	fmt.Println("UCS2 positions with characters:")
	for _, c := range smslite.UCS2CharacterPositions(text) {
		fmt.Printf("  rune=%d byte=%d utf16=%d char=%q visible=%s code=%s\n", c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint)
	}
	fmt.Println(smslite.RenderTextTable(a))
}
