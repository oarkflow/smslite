package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oarkflow/smslite"
)

func main() {
	jsonOut := flag.Bool("json", false, "print JSON report")
	htmlOut := flag.Bool("html", false, "print HTML report")
	ucs2Out := flag.Bool("ucs2", false, "print only UCS-2 forcing character positions with characters")
	invisOut := flag.Bool("invisible", false, "print invisible/zero-width/bidi/format character positions with characters")
	normalize := flag.Bool("normalize", false, "print GSM-7 normalization suggestion")
	clean := flag.Bool("clean", false, "strictly clean pasted text to GSM-7 by mapping known Unicode/confusable characters and dropping unknown non-GSM characters")
	flag.Parse()
	text := strings.Join(flag.Args(), " ")
	if text == "" {
		buf, _ := os.ReadFile("/dev/stdin")
		text = strings.TrimRight(string(buf), "\n")
	}
	if *normalize {
		n := smslite.NormalizeForGSM7(text)
		fmt.Println(n.Normalized)
		return
	}
	if *clean {
		c := smslite.CleanForGSM7(text)
		fmt.Println(c.Cleaned)
		return
	}
	if *invisOut {
		items := smslite.InvisibleCharacterPositions(text)
		if len(items) == 0 {
			fmt.Println("no invisible, zero-width, bidi, format-control, or Unicode-space characters found")
			return
		}
		fmt.Println("rune\tbyte\tutf16\tchar\tvisible\tcode\tclass\tname\treason")
		for _, c := range items {
			fmt.Printf("%d\t%d\t%d\t%s\t%s\t%s\t%s\t%s\t%s\n", c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint, c.ClassLabel, c.Name, c.Reason)
		}
		return
	}
	if *ucs2Out {
		items := smslite.UCS2CharacterPositions(text)
		if len(items) == 0 {
			fmt.Println("no UCS-2 forcing characters; text is GSM-7 compatible")
			return
		}
		fmt.Println("rune\tbyte\tutf16\tchar\tvisible\tcode\treason")
		for _, c := range items {
			fmt.Printf("%d\t%d\t%d\t%s\t%s\t%s\t%s\n", c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint, c.Reason)
		}
		return
	}
	a := smslite.AnalyzeDetailed(text)
	if *jsonOut {
		b, _ := smslite.ToJSONReport(a)
		os.Stdout.Write(b)
		os.Stdout.Write([]byte("\n"))
		return
	}
	if *htmlOut {
		fmt.Println(smslite.RenderHTML(a))
		return
	}
	fmt.Println(smslite.FormatAnalysisSummary(a))
	ucs2 := smslite.UCS2CharacterPositions(text)
	if len(ucs2) > 0 {
		fmt.Println("UCS-2 forcing characters:")
		for _, c := range ucs2 {
			fmt.Printf("  rune=%d byte=%d utf16=%d char=%q visible=%s code=%s\n", c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint)
		}
	}
	fmt.Print(smslite.RenderTextTable(a))
}
