# CLI examples

```bash
go run ./cmd/smsanalyze "Hello {name}"
go run ./cmd/smsanalyze -ucs2 "Hello नमस्ते 🙂"
go run ./cmd/smsanalyze -invisible $'A\u2060B\u2066C\u2069'
go run ./cmd/smsanalyze -json "Hello 🙂"
go run ./cmd/smsanalyze -html "Hello 🙂" > report.html
go run ./cmd/smsanalyze -normalize "Hello “world”—ok…"
```
