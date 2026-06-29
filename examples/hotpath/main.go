package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "Hello {name}, use code 123456"
	s := smslite.Analyze(text) // zero allocations; message-level facts only
	fmt.Println(smslite.FormatSummary(s))
	fmt.Printf("encoding=%s units=%d parts=%d multipart=%v\n", s.Encoding, s.LengthUnits, s.SegmentCount, s.IsMultipart)
}
