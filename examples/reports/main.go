package main

import (
	"fmt"
	"os"

	"github.com/oarkflow/smslite"
)

func main() {
	a := smslite.AnalyzeDetailed("Pay\u2066now\u2069 🙂")
	json := smslite.RenderJSON(a)
	fmt.Println(string(json))
	_ = os.WriteFile("smslite-report.html", []byte(smslite.RenderHTML(a)), 0o644)
}
