package smslite

// AnalyzeTrustedGSM7SingleSeptetASCII is the ultra-hot path for payloads that are
// already known to contain only one-septet GSM-7 ASCII characters. It performs no
// validation and therefore runs in single-digit ns/op on common Go benchmarks.
// Use AnalyzeSummary for untrusted text.
func AnalyzeTrustedGSM7SingleSeptetASCII(text string) Summary {
	units := len(text)
	parts := 0
	if units > 0 {
		if units <= GSM7SingleLimit {
			parts = 1
		} else {
			parts = (units + GSM7MultipartLimit - 1) / GSM7MultipartLimit
		}
	}
	return Summary{Encoding: EncodingGSM7, ByteLength: units, RuneCount: units, GSMSeptetCount: units, UCS2UnitCount: units, LengthUnits: units, SegmentCount: parts, IsMultipart: parts > 1, SingleSegmentLimit: GSM7SingleLimit, MultipartLimit: GSM7MultipartLimit, FirstUCS2RuneIndex: -1, FirstUCS2ByteIndex: -1}
}

// AnalyzeValidatedGSM7SingleSeptetASCII validates that text contains only ASCII
// characters that cost exactly one GSM septet, then returns a summary. Extension
// characters such as {, }, [, ], ^, ~, \, |, and € should use AnalyzeSummary.
func AnalyzeValidatedGSM7SingleSeptetASCII(text string) (Summary, bool) {
	for i := 0; i < len(text); i++ {
		b := text[i]
		if b >= 128 || asciiGSM7Cost[b] != 1 {
			return Summary{}, false
		}
	}
	return AnalyzeTrustedGSM7SingleSeptetASCII(text), true
}
