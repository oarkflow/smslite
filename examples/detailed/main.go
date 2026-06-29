package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "Hi ⁠there 🙂"
	a := smslite.AnalyzeDetailed(text)
	fmt.Println(smslite.FormatAnalysisSummary(a))
	fmt.Println(smslite.RenderTextTable(a))
}
