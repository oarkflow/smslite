package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	text := "Hello “world”—use code…"
	res := smslite.NormalizeForGSM7(text)
	fmt.Println("before:", text)
	fmt.Println("after: ", res.Normalized)
	for _, s := range res.Suggestions {
		fmt.Printf("rune=%d %q -> %q reason=%s\n", s.IndexRune, s.Original, s.Replacement, s.Reason)
	}
}
