package smslite

import "testing"

func TestEncoding(t *testing.T) {
	if DetectEncoding("Hello {}[]^~\\|€") != EncodingGSM7 {
		t.Fatal("expected GSM7")
	}
	if DetectEncoding("Hello 🙂") != EncodingUCS2 {
		t.Fatal("expected UCS2")
	}
}

func TestLengths(t *testing.T) {
	n, err := GSM7SeptetLength("{}")
	if err != nil || n != 4 {
		t.Fatalf("septets=%d err=%v", n, err)
	}
	if UCS2UnitLength("🙂") != 2 {
		t.Fatal("emoji should be 2 UTF-16 units")
	}
}

func TestSegmentsGSM(t *testing.T) {
	msg := makeString('a', 161)
	s := CalculateSegmentsDetailed(msg)
	if s.Encoding != EncodingGSM7 || s.SegmentCount != 2 || !s.IsMultipart {
		t.Fatalf("bad seg: %+v", s)
	}
	if len(s.Segments) != 2 || s.Segments[0].LengthUnits != 153 || s.Segments[1].LengthUnits != 8 {
		t.Fatalf("bad parts: %+v", s.Segments)
	}
}

func TestSegmentsUCS2(t *testing.T) {
	msg := makeString('न', 71)
	s := CalculateSegmentsDetailed(msg)
	if s.Encoding != EncodingUCS2 || s.SegmentCount != 2 {
		t.Fatalf("bad seg: %+v", s)
	}
}

func TestUCS2Positions(t *testing.T) {
	p := UCS2Positions("Hi 🙂 नमस्ते")
	if len(p) == 0 || p[0] != 3 {
		t.Fatalf("bad positions: %#v", p)
	}
}

func TestUDH(t *testing.T) {
	u := BuildConcatUDH8(0xAB, 2, 1)
	if u.Hex != "05 00 03 AB 02 01" {
		t.Fatal(u.Hex)
	}
	p, err := ParseUDH([]byte{0x05, 0x00, 0x03, 0xAB, 0x02, 0x01})
	if err != nil || p.ReferenceNumber != 0xAB || p.TotalParts != 2 || p.PartNumber != 1 {
		t.Fatalf("%+v %v", p, err)
	}
}

func TestNormalize(t *testing.T) {
	n := NormalizeForGSM7("Hello “world”—ok…")
	if !n.Changed || !IsGSM7(n.Normalized) {
		t.Fatalf("bad normalize: %+v", n)
	}
}

func makeString(r rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = r
	}
	return string(b)
}

func BenchmarkDetectGSM7(b *testing.B) {
	s := "Hello World 1234567890"
	for i := 0; i < b.N; i++ {
		_ = DetectEncoding(s)
	}
}
func BenchmarkGSM7SeptetLengthFast(b *testing.B) {
	s := "Hello {name} []^~\\|€"
	for i := 0; i < b.N; i++ {
		_, _ = GSM7SeptetLengthFast(s)
	}
}
func BenchmarkUCS2UnitLength(b *testing.B) {
	s := "Hello नमस्ते 🙂"
	for i := 0; i < b.N; i++ {
		_ = UCS2UnitLength(s)
	}
}
func BenchmarkAnalyze(b *testing.B) {
	s := "Hello {name}, use code 123456"
	for i := 0; i < b.N; i++ {
		_ = Analyze(s)
	}
}
func BenchmarkAnalyzeSummary(b *testing.B) {
	s := "Hello {name}, use code 123456"
	for i := 0; i < b.N; i++ {
		_ = AnalyzeSummary(s)
	}
}

func TestUCS2CharacterPositionsIncludeCharacters(t *testing.T) {
	items := UCS2CharacterPositions("Hi 🙂 नमस्ते")
	if len(items) == 0 {
		t.Fatal("expected UCS2 character details")
	}
	if items[0].IndexRune != 3 || items[0].Char != "🙂" || items[0].CodePoint != "U+1F642" {
		t.Fatalf("bad first item: %+v", items[0])
	}
	first := FirstUCS2CharacterPosition("Hello—world")
	if first == nil || first.Char != "—" || first.Visible != "—" {
		t.Fatalf("bad first position: %+v", first)
	}
}

func TestTrustedASCII(t *testing.T) {
	s, ok := AnalyzeValidatedGSM7SingleSeptetASCII("Hello 123")
	if !ok || s.Encoding != EncodingGSM7 || s.LengthUnits != 9 {
		t.Fatalf("bad validated summary: %+v ok=%v", s, ok)
	}
	if _, ok := AnalyzeValidatedGSM7SingleSeptetASCII("Hello {name}"); ok {
		t.Fatal("extension chars are not single-septet ASCII")
	}
}

func BenchmarkAnalyzeTrustedGSM7SingleSeptetASCII(b *testing.B) {
	s := "Hello name use code 123456"
	for i := 0; i < b.N; i++ {
		_ = AnalyzeTrustedGSM7SingleSeptetASCII(s)
	}
}

func BenchmarkAnalyzeValidatedGSM7SingleSeptetASCII(b *testing.B) {
	s := "Hello name use code 123456"
	for i := 0; i < b.N; i++ {
		_, _ = AnalyzeValidatedGSM7SingleSeptetASCII(s)
	}
}

func TestComprehensiveInvisibleCharacters(t *testing.T) {
	text := "A\u205F\u2060\u2061\u2062\u2063\u2064\u2065\u2066\u2067\u2068\u2069\u206A\u206B\u206C\u206D\u206E\u206FB"
	items := InvisibleCharacterPositions(text)
	if got, want := len(items), 17; got != want {
		t.Fatalf("invisible count=%d want=%d: %#v", got, want, items)
	}
	checks := map[rune]string{
		0x205F: "MEDIUM MATHEMATICAL SPACE",
		0x2060: "WORD JOINER",
		0x2061: "FUNCTION APPLICATION",
		0x2062: "INVISIBLE TIMES",
		0x2063: "INVISIBLE SEPARATOR",
		0x2064: "INVISIBLE PLUS",
		0x2065: "INHIBIT SYMMETRIC SWAPPING",
		0x2066: "LEFT-TO-RIGHT ISOLATE",
		0x2067: "RIGHT-TO-LEFT ISOLATE",
		0x2068: "FIRST STRONG ISOLATE",
		0x2069: "POP DIRECTIONAL ISOLATE",
		0x206A: "INHIBIT SYMMETRIC SWAPPING",
		0x206B: "ACTIVATE SYMMETRIC SWAPPING",
		0x206C: "INHIBIT ARABIC FORM SHAPING",
		0x206D: "ACTIVATE ARABIC FORM SHAPING",
		0x206E: "NATIONAL DIGIT SHAPES",
		0x206F: "NOMINAL DIGIT SHAPES",
	}
	for _, item := range items {
		r := []rune(item.Char)[0]
		if want := checks[r]; want != "" && item.Name != want {
			t.Fatalf("%U name=%q want=%q", r, item.Name, want)
		}
	}
	for _, r := range []rune{0x205F, 0x2060, 0x2061, 0x2062, 0x2063, 0x2064, 0x2065, 0x2066, 0x2067, 0x2068, 0x2069, 0x206A, 0x206B, 0x206C, 0x206D, 0x206E, 0x206F} {
		if !IsInvisible(r) {
			t.Fatalf("%U should be invisible", r)
		}
		if VisibleChar(r) == string(r) {
			t.Fatalf("%U should have visible replacement", r)
		}
	}
}

func TestInvisibleCharactersInAnalysis(t *testing.T) {
	a := AnalyzeDetailed("pay\u2066now\u2069")
	if a.Encoding != EncodingUCS2 {
		t.Fatalf("encoding=%s want UCS-2", a.Encoding)
	}
	if len(a.InvisibleCharacters) != 2 {
		t.Fatalf("invisible=%d want 2", len(a.InvisibleCharacters))
	}
	if a.InvisibleCharacters[0].Visible != "⟦LRI⟧" || !a.Characters[3].IsBidiControl {
		t.Fatalf("bidi details missing: %#v %#v", a.InvisibleCharacters[0], a.Characters[3])
	}
}

func TestCleanTextToGSMRobustMappings(t *testing.T) {
	in := "ΡΑΥΡΑL раураl H\u200Be\u200Cl\u200Dl\u2060o 𝓗𝓮𝓵𝓵𝓸 Ｂｉｇ café “ok”—₹🙂"
	got := CleanTextToGSM(in)
	want := "PAYPAL paypal Hello Hello Big cafe \"ok\"-Rs"
	if got != want {
		t.Fatalf("cleaned=%q want %q", got, want)
	}
	if !IsGSM7(got) {
		t.Fatalf("cleaned text must be GSM-7: %q", got)
	}
}

func TestCleanForGSM7ReportsMappingAndRemoval(t *testing.T) {
	res := CleanForGSM7("A\u200BB🙂é")
	if res.Cleaned != "ABe" || res.Encoding != EncodingGSM7 || !res.Changed {
		t.Fatalf("bad clean result: %+v", res)
	}
	if res.Mapped == 0 || res.Removed != 2 {
		t.Fatalf("bad counts: mapped=%d removed=%d replacements=%+v", res.Mapped, res.Removed, res.Replacements)
	}
}

func TestCleanToGSMUserSampleMappings(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "zero width removed", in: "Hel\u200Blo\u200D World\uFEFF", want: "Hello World"},
		{name: "smart punctuation", in: "“Hello”—it’s OK…", want: "\"Hello\"-it's OK..."},
		{name: "greek and cyrillic spoof letters", in: "Ρаураl ΑΒС", want: "Paypal ABC"},
		{name: "fullwidth", in: "Ｈｅｌｌｏ １２３！", want: "Hello 123!"},
		{name: "math styled", in: "𝐇𝐞𝐥𝐥𝐨 𝟙𝟚𝟛", want: "Hello 123"},
		{name: "accented latin", in: "Café Noël São Paulo", want: "Cafe Noel Sao Paulo"},
		{name: "unsupported removed", in: "OTP is 123456 ✅ 🔥", want: "OTP is 123456  "},
	}

	for _, tt := range tests {
		got := CleanToGSM(tt.in)
		if got != tt.want {
			t.Fatalf("%s: got %q, want %q", tt.name, got, tt.want)
		}
		if !IsGSM7(got) {
			t.Fatalf("%s: cleaned value must be GSM-7: %q", tt.name, got)
		}
	}
}

func TestCleanToGSMExtraDigitAndInvisibleCoverage(t *testing.T) {
	in := "OTP १२३ ۱۲۳ ١٢٣\U0001F1F3\U0001F1F5\U000E0001\u034Fµ Чч"
	want := "OTP 123 123 123NPu Yy"
	got := CleanToGSM(in)
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
	if RequiresUCS2(got) {
		t.Fatalf("cleaned output must not require UCS-2: %q", got)
	}
}
