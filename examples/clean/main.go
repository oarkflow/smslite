package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "ΡΑΥΡΑL раураl H\u200Be\u200Cl\u200Dl\u2060o 𝓗𝓮𝓵𝓵𝓸 Ｂｉｇ café “ok”—₹🙂"
	result := smslite.CleanForGSM7(text)
	fmt.Println(result.Cleaned)
	fmt.Println(result.Encoding)
	fmt.Printf("mapped=%d removed=%d\n", result.Mapped, result.Removed)
}
