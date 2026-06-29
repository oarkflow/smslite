package main

import (
	"fmt"
	"strings"

	"github.com/oarkflow/smslite"
)

func main() {
	gsm := strings.Repeat("a", 161)
	seg := smslite.CalculateSegmentsDetailed(gsm)
	fmt.Printf("GSM encoding=%s units=%d parts=%d multipart=%v\n", seg.Encoding, seg.LengthUnits, seg.SegmentCount, seg.IsMultipart)
	for _, p := range seg.Segments {
		fmt.Printf("part=%d/%d units=%d bytes=%d:%d udh=%s\n", p.PartNumber, p.TotalParts, p.LengthUnits, p.StartByteIndex, p.EndByteIndex, p.UDH.Hex)
	}

	ucs := strings.Repeat("न", 71)
	seg = smslite.CalculateSegmentsDetailed(ucs)
	fmt.Printf("UCS2 encoding=%s units=%d parts=%d multipart=%v\n", seg.Encoding, seg.LengthUnits, seg.SegmentCount, seg.IsMultipart)
}
