package smslite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"strings"
)

type CharacterVisualization struct {
	Text       string          `json:"text"`
	Encoding   Encoding        `json:"encoding"`
	Characters []CharacterInfo `json:"characters"`
	Table      string          `json:"table"`
	HTML       string          `json:"html"`
}

func VisualizeCharacters(text string) CharacterVisualization {
	a := AnalyzeDetailed(text)
	return CharacterVisualization{Text: text, Encoding: a.Encoding, Characters: a.Characters, Table: RenderTextTable(a), HTML: RenderHTML(a)}
}

func ToJSONReport(a SMSAnalysis) ([]byte, error) { return json.MarshalIndent(a, "", "  ") }
func RenderJSON(a SMSAnalysis) []byte            { b, _ := ToJSONReport(a); return b }

func RenderTextTable(a SMSAnalysis) string {
	var b bytes.Buffer
	b.WriteString("#\tChar\tCode\tName\tKind\tSeptets\tUCS2\tSegment\n")
	for _, c := range a.Characters {
		b.WriteString(strconv.Itoa(c.IndexRune))
		b.WriteByte('\t')
		b.WriteString(c.Visible)
		b.WriteByte('\t')
		b.WriteString(c.CodePoint)
		b.WriteByte('\t')
		if c.UnicodeName != "" {
			b.WriteString(c.UnicodeName)
		} else {
			b.WriteByte('-')
		}
		b.WriteByte('\t')
		b.WriteString(c.KindLabel)
		b.WriteByte('\t')
		b.WriteString(strconv.Itoa(c.GSMSeptets))
		b.WriteByte('\t')
		b.WriteString(strconv.Itoa(c.UCS2Units))
		b.WriteByte('\t')
		b.WriteString(strconv.Itoa(c.SegmentIndex + 1))
		b.WriteByte('\n')
	}
	return b.String()
}

func RenderHTML(a SMSAnalysis) string {
	var b strings.Builder
	b.Grow(len(a.Characters) * 128)
	b.WriteString(`<div class="smslite-report">`)
	b.WriteString(`<div class="smslite-summary">Encoding: <b>` + a.Encoding.String() + `</b>, Segments: <b>` + strconv.Itoa(a.SegmentCount) + `</b>, Units: <b>` + strconv.Itoa(a.LengthUnits) + `</b></div>`)
	b.WriteString(`<table class="smslite-chars"><thead><tr><th>#</th><th>Char</th><th>Code</th><th>Name</th><th>Kind</th><th>Septets</th><th>UCS2</th><th>Segment</th></tr></thead><tbody>`)
	for _, c := range a.Characters {
		cls := "gsm-default"
		if c.IsGSMExtended {
			cls = "gsm-extended"
		}
		if c.IsUCS2 {
			cls = "ucs2"
		}
		if c.IsEmoji {
			cls = "emoji"
		}
		if c.IsInvisible {
			cls += " invisible"
		}
		b.WriteString(`<tr class="` + cls + `"><td>` + strconv.Itoa(c.IndexRune) + `</td><td title="` + html.EscapeString(c.Char) + `">` + html.EscapeString(c.Visible) + `</td><td>` + c.CodePoint + `</td><td>` + html.EscapeString(c.UnicodeName) + `</td><td>` + html.EscapeString(c.KindLabel) + `</td><td>` + strconv.Itoa(c.GSMSeptets) + `</td><td>` + strconv.Itoa(c.UCS2Units) + `</td><td>` + strconv.Itoa(c.SegmentIndex+1) + `</td></tr>`)
	}
	b.WriteString(`</tbody></table></div>`)
	return b.String()
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("encoding=%s units=%d segments=%d multipart=%t gsmSeptets=%d ucs2Units=%d firstUCS2Rune=%d firstUCS2Byte=%d", s.Encoding, s.LengthUnits, s.SegmentCount, s.IsMultipart, s.GSMSeptetCount, s.UCS2UnitCount, s.FirstUCS2RuneIndex, s.FirstUCS2ByteIndex)
}

func FormatAnalysisSummary(a SMSAnalysis) string {
	return fmt.Sprintf("encoding=%s units=%d segments=%d multipart=%t gsmSeptets=%d ucs2Units=%d", a.Encoding, a.LengthUnits, a.SegmentCount, a.IsMultipart, a.GSMSeptetCount, a.UCS2UnitCount)
}
