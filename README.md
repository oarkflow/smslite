# smslite

High-performance, lightweight SMS encoding, segmentation, invisible-character, and visual-analysis toolkit for Go.

The default `Analyze(text)` path is designed for production sending pipelines: it performs message-wide GSM-7/UCS-2 detection, length calculation, multipart calculation, and first UCS-2 position reporting with `0 B/op` and `0 allocs/op`.

Full visual diagnostics are available through `AnalyzeDetailed(text)` and report renderers. Diagnostic/reporting APIs intentionally return slices/strings/JSON/HTML, so they are separated from the allocation-free hot path.

## Features

- GSM-7 default alphabet detection.
- GSM-7 extension table detection with 2-septet accounting for `^ { } \\ [ ~ ] | €`.
- Message-wide UCS-2 fallback detection.
- Exact UCS-2 forcing character positions with character, visible label, code point, byte index, rune index, UTF-16 index, and reason.
- UTF-8 byte index, rune index, and UTF-16/UCS-2 unit index tracking.
- Emoji, combining mark, variation selector, control, zero-width, bidi, math invisible, format-control, and Unicode-space detection.
- SMS segment calculation using GSM-7 `160 / 153` and UCS-2 `70 / 67` limits.
- Multipart detection with exact total parts.
- Optional segment splitting with per-part text boundaries and UDH metadata.
- 8-bit and 16-bit concatenated SMS UDH build/parse.
- Text, JSON, and HTML character reports.
- Smart punctuation normalization suggestions.
- Billing helpers.
- CLI: `smsanalyze`.
- Examples for hot path, detailed reports, UCS2 positions, invisible characters, multipart, UDH, normalization, reports, and CLI.

## Install

```bash
go test ./...
go run ./cmd/smsanalyze "Hello {name}"
```

## CLI

```bash
go run ./cmd/smsanalyze "Hello {name} 🙂"
go run ./cmd/smsanalyze -json "Hello नमस्ते 🙂"
go run ./cmd/smsanalyze -ucs2 "Hello नमस्ते 🙂"
go run ./cmd/smsanalyze -invisible $'A\u2060B\u2066C\u2069'
go run ./cmd/smsanalyze -html "Hello World" > report.html
echo 'Hello “smart quotes”' | go run ./cmd/smsanalyze -normalize
```

## Hot path API: zero allocations

```go
s := smslite.Analyze("Hello {name}, use code 123456")
fmt.Println(s.Encoding, s.LengthUnits, s.SegmentCount, s.IsMultipart)
```

`Analyze(text)` returns `Summary` and does not allocate. It includes:

- encoding
- byte length
- rune count
- GSM septet count
- UCS-2/UTF-16 unit count
- effective billing units
- segment count
- multipart flag
- single/multipart limits
- first UCS-2 rune and byte position

## Detailed analysis API

```go
a := smslite.AnalyzeDetailed("Hi ⁠there 🙂")
fmt.Println(smslite.FormatAnalysisSummary(a))
fmt.Println(smslite.RenderTextTable(a))
```

Use this when you need full diagnostic slices:

- `Characters []CharacterInfo`
- `UCS2Characters []CharacterInfo`
- `InvisibleCharacters []InvisibleCharacterPosition`
- `Warnings []Warning`
- `Recommendations []Suggestion`
- `Segments []SMSSegment`

## Essential APIs

### Encoding and summary

- `Analyze(text string) Summary`
- `AnalyzeSummary(text string) Summary`
- `AnalyzeSummaryASCII(text string) Summary`
- `AnalyzeTrustedGSM7SingleSeptetASCII(text string) Summary`
- `AnalyzeValidatedGSM7SingleSeptetASCII(text string) (Summary, bool)`
- `DetectEncoding(text string) Encoding`
- `DetectEncodingDetailed(text string) EncodingResult`
- `IsGSM7(text string) bool`
- `IsGSM7Default(r rune) bool`
- `IsGSM7Extended(r rune) bool`
- `IsUCS2Required(text string) bool`

### Detailed analysis and reporting

- `AnalyzeDetailed(text string) SMSAnalysis`
- `AnalyzeDetailedWithOptions(text string, opt Options) SMSAnalysis`
- `AnalyzeWithOptions(text string, opt Options) SMSAnalysis`
- `VisualizeCharacters(text string) CharacterVisualization`
- `RenderTextTable(analysis SMSAnalysis) string`
- `RenderHTML(analysis SMSAnalysis) string`
- `ToJSONReport(analysis SMSAnalysis) ([]byte, error)`
- `RenderJSON(analysis SMSAnalysis) []byte`

### UCS-2 positions with actual characters

- `UCS2Positions(text string) []int`
- `UCS2CharacterPositions(text string) []UCS2CharacterPosition`
- `AppendUCS2CharacterPositions(dst []UCS2CharacterPosition, text string) []UCS2CharacterPosition`
- `FirstUCS2CharacterPosition(text string) *UCS2CharacterPosition`
- `FindUCS2Characters(text string) []CharacterInfo`
- `FirstUCS2Character(text string) *CharacterInfo`

Example:

```go
for _, c := range smslite.UCS2CharacterPositions("Hello नमस्ते 🙂") {
    fmt.Printf("rune=%d byte=%d utf16=%d char=%q visible=%s code=%s\n",
        c.IndexRune, c.IndexByte, c.IndexUTF16, c.Char, c.Visible, c.CodePoint)
}
```

### Invisible, zero-width, bidi, and format controls

- `InvisibleCharacterPositions(text string) []InvisibleCharacterPosition`
- `AppendInvisibleCharacterPositions(dst []InvisibleCharacterPosition, text string) []InvisibleCharacterPosition`
- `FirstInvisibleCharacterPosition(text string) *InvisibleCharacterPosition`
- `FindInvisibleCharacters(text string) []CharacterInfo`
- `IsInvisible(r rune) bool`
- `IsZeroWidth(r rune) bool`
- `IsBidiControl(r rune) bool`
- `IsFormatControl(r rune) bool`
- `IsUnicodeSpaceSeparator(r rune) bool`
- `IsMathInvisible(r rune) bool`
- `VisibleChar(r rune) string`

### Length, segmentation, and UDH

- `GSM7SeptetLength(text string) (int, error)`
- `GSM7SeptetLengthFast(text string) (int, bool)`
- `UCS2UnitLength(text string) int`
- `UCS2Units(r rune) int`
- `CalculateSegments(text string) Segmentation` allocation-free summary segmentation, no segment slice
- `CalculateSegmentsDetailed(text string) Segmentation` includes `Segments []SMSSegment`
- `SplitSMS(text string) ([]SMSSegment, error)`
- `SplitGSM7(text string) ([]SMSSegment, error)`
- `SplitUCS2(text string) ([]SMSSegment, error)`
- `BuildConcatUDH8(ref byte, total, part byte) UDHInfo`
- `BuildConcatUDH16(ref uint16, total, part byte) UDHInfo`
- `ParseUDH(data []byte) (UDHInfo, error)`

### Normalization and billing

- `NormalizeForGSM7(text string) NormalizationResult`
- `SuggestReplacements(text string) []Suggestion`
- `EstimateBillingUnits(text string) BillingInfo`
- `RemainingUnits(text string) int`
- `RemainingCharacters(text string) int`
- `NextCharacterCost(r rune, currentText string) int`

## Comprehensive invisible / zero-width / bidi detection

`smslite` reports the actual hidden character, visible replacement label, byte/rune/UTF-16 position, Unicode code point, name, class, and reason.

Covered classes include:

- ASCII controls and whitespace.
- NBSP and Unicode space separators: `U+00A0`, `U+1680`, `U+2000`-`U+200A`, `U+2028`, `U+2029`, `U+202F`, `U+205F`, `U+3000`.
- Zero-width characters: `U+200B`, `U+200C`, `U+200D`, `U+2060`, `U+FEFF`.
- Mathematical invisible operators: `U+2061`, `U+2062`, `U+2063`, `U+2064`.
- Bidirectional marks, embeddings, overrides, and isolates: `U+200E`, `U+200F`, `U+202A`-`U+202E`, `U+2066`-`U+2069`.
- Deprecated format controls: `U+180E`, `U+2065`, `U+206A`-`U+206F`.
- Variation selectors: `U+FE00`-`U+FE0F`, `U+E0100`-`U+E01EF`.
- Combining marks are visually represented as `◌` + mark in full character reports.

Example:

```bash
go run ./cmd/smsanalyze -invisible $'Hello\u2066world\u2069'
```

Output includes:

```text
rune byte utf16 char visible code class name reason
5    5    5     ⁦    ⟦LRI⟧ U+2066 bidirectional control LEFT-TO-RIGHT ISOLATE ...
```

## Performance

Run:

```bash
go test -bench . -benchmem
```

Current benchmark in this environment:

```text
BenchmarkDetectGSM7-56                                53.66 ns/op    0 B/op    0 allocs/op
BenchmarkGSM7SeptetLengthFast-56                      36.78 ns/op    0 B/op    0 allocs/op
BenchmarkUCS2UnitLength-56                            63.32 ns/op    0 B/op    0 allocs/op
BenchmarkAnalyze-56                                   41.37 ns/op    0 B/op    0 allocs/op
BenchmarkAnalyzeSummary-56                            41.22 ns/op    0 B/op    0 allocs/op
BenchmarkAnalyzeTrustedGSM7SingleSeptetASCII-56        0.46 ns/op    0 B/op    0 allocs/op
BenchmarkAnalyzeValidatedGSM7SingleSeptetASCII-56     32.15 ns/op    0 B/op    0 allocs/op
```

On faster desktop CPUs, the same hot-path benchmarks are typically lower. The hot-path APIs are the ones to use inside high-throughput SMS routing and sending services.

## Examples

```bash
go run ./examples/hotpath
go run ./examples/detailed
go run ./examples/ucs2_positions
go run ./examples/invisible
go run ./examples/multipart
go run ./examples/udh
go run ./examples/normalize
go run ./examples/reports
```

## Allocation model

`Analyze`, `AnalyzeSummary`, `DetectEncoding`, `GSM7SeptetLengthFast`, `UCS2UnitLength`, and the trusted/validated ASCII paths are the no-allocation runtime APIs.

Full visual reports allocate because Go slices and rendered JSON/HTML/table strings must be materialized. For lower allocation diagnostic flows, reuse caller-owned buffers with append-style APIs such as:

```go
ucs2Buf := make([]smslite.UCS2CharacterPosition, 0, 32)
ucs2Buf = smslite.AppendUCS2CharacterPositions(ucs2Buf[:0], text)

invBuf := make([]smslite.InvisibleCharacterPosition, 0, 32)
invBuf = smslite.AppendInvisibleCharacterPositions(invBuf[:0], text)
```
