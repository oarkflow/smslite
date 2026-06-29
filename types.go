package smslite

import "encoding/json"

type Encoding uint8

const (
	EncodingGSM7 Encoding = iota
	EncodingUCS2
	EncodingBinary
)

func (e Encoding) String() string {
	switch e {
	case EncodingGSM7:
		return "GSM-7"
	case EncodingUCS2:
		return "UCS-2"
	case EncodingBinary:
		return "BINARY"
	default:
		return "UNKNOWN"
	}
}

func (e Encoding) MarshalJSON() ([]byte, error) { return json.Marshal(e.String()) }

type DetectionMode uint8

const (
	DetectionStrict DetectionMode = iota
	DetectionRelaxed
)

type Options struct {
	Mode                 DetectionMode
	GraphemeSafe         bool
	PreserveWords        bool
	TrimSegmentEdges     bool
	ReferenceNumberSize  int
	ReferenceNumber      uint16
	MaxParts             int
	CollectCharacters    bool
	CollectWarnings      bool
	CollectSegments      bool
	CollectVisualization bool
}

func DefaultOptions() Options {
	return Options{Mode: DetectionStrict, ReferenceNumberSize: 8, ReferenceNumber: 0xAB, MaxParts: 255}
}

// DiagnosticOptions enables full report collection. These APIs allocate because
// they return slices and strings.
func DiagnosticOptions() Options {
	return Options{Mode: DetectionStrict, ReferenceNumberSize: 8, ReferenceNumber: 0xAB, MaxParts: 255, CollectCharacters: true, CollectWarnings: true, CollectSegments: true, CollectVisualization: true}
}

type SMSAnalysis struct {
	Text                string                       `json:"text,omitempty"`
	Encoding            Encoding                     `json:"encoding"`
	EncodingReason      string                       `json:"encodingReason"`
	ByteLength          int                          `json:"byteLength"`
	RuneCount           int                          `json:"runeCount"`
	GSMSeptetCount      int                          `json:"gsmSeptetCount"`
	UCS2UnitCount       int                          `json:"ucs2UnitCount"`
	LengthUnits         int                          `json:"lengthUnits"`
	SegmentCount        int                          `json:"segmentCount"`
	MaxCharsPerSegment  int                          `json:"maxCharsPerSegment"`
	SingleSegmentLimit  int                          `json:"singleSegmentLimit"`
	MultipartLimit      int                          `json:"multipartLimit"`
	IsMultipart         bool                         `json:"isMultipart"`
	Segments            []SMSSegment                 `json:"segments,omitempty"`
	Characters          []CharacterInfo              `json:"characters,omitempty"`
	UCS2Characters      []CharacterInfo              `json:"ucs2Characters,omitempty"`
	InvisibleCharacters []InvisibleCharacterPosition `json:"invisibleCharacters,omitempty"`
	Warnings            []Warning                    `json:"warnings,omitempty"`
	Recommendations     []Suggestion                 `json:"recommendations,omitempty"`
	FirstUCS2RuneIndex  int                          `json:"firstUCS2RuneIndex"`
	FirstUCS2ByteIndex  int                          `json:"firstUCS2ByteIndex"`
}

type CharacterKind uint8

const (
	KindGSMDefault CharacterKind = iota
	KindGSMExtended
	KindUCS2
	KindControl
	KindInvisible
	KindCombining
	KindEmoji
)

func (k CharacterKind) String() string {
	switch k {
	case KindGSMDefault:
		return "GSM default"
	case KindGSMExtended:
		return "GSM extension"
	case KindUCS2:
		return "UCS-2"
	case KindControl:
		return "control"
	case KindInvisible:
		return "invisible"
	case KindCombining:
		return "combining"
	case KindEmoji:
		return "emoji"
	default:
		return "unknown"
	}
}

// UCS2CharacterPosition is a compact allocation-light record for characters that force
// message-wide UCS-2 encoding. It is smaller than CharacterInfo and intended for
// position reports, CLI output, logs, and validation errors.
type UCS2CharacterPosition struct {
	IndexByte  int    `json:"indexByte"`
	IndexRune  int    `json:"indexRune"`
	IndexUTF16 int    `json:"indexUTF16"`
	Char       string `json:"char"`
	Visible    string `json:"visible"`
	CodePoint  string `json:"codePoint"`
	Hex        string `json:"hex"`
	Reason     string `json:"reason"`
}

type CharacterInfo struct {
	IndexByte           int            `json:"indexByte"`
	IndexRune           int            `json:"indexRune"`
	IndexUTF16          int            `json:"indexUTF16"`
	Char                string         `json:"char"`
	Rune                rune           `json:"rune"`
	CodePoint           string         `json:"codePoint"`
	UnicodeName         string         `json:"unicodeName,omitempty"`
	Category            string         `json:"category,omitempty"`
	Kind                CharacterKind  `json:"kind"`
	KindLabel           string         `json:"kindLabel"`
	Encoding            Encoding       `json:"encoding"`
	IsGSMDefault        bool           `json:"isGSMDefault"`
	IsGSMExtended       bool           `json:"isGSMExtended"`
	IsUCS2              bool           `json:"isUCS2"`
	IsEmoji             bool           `json:"isEmoji"`
	IsCombiningMark     bool           `json:"isCombiningMark"`
	IsVariation         bool           `json:"isVariation"`
	IsZeroWidthJoiner   bool           `json:"isZeroWidthJoiner"`
	IsZeroWidth         bool           `json:"isZeroWidth"`
	IsBidiControl       bool           `json:"isBidiControl"`
	IsFormatControl     bool           `json:"isFormatControl"`
	IsMathInvisible     bool           `json:"isMathInvisible"`
	InvisibleClass      InvisibleClass `json:"invisibleClass"`
	InvisibleClassLabel string         `json:"invisibleClassLabel,omitempty"`
	IsInvisible         bool           `json:"isInvisible"`
	GSMSeptets          int            `json:"gsmSeptets"`
	UCS2Units           int            `json:"ucs2Units"`
	SegmentIndex        int            `json:"segmentIndex"`
	PositionInSegment   int            `json:"positionInSegment"`
	Visible             string         `json:"visible"`
	DisplayLabel        string         `json:"displayLabel"`
	Hex                 string         `json:"hex"`
	Warning             string         `json:"warning,omitempty"`
}

type SMSSegment struct {
	Index          int             `json:"index"`
	PartNumber     int             `json:"partNumber"`
	TotalParts     int             `json:"totalParts"`
	Text           string          `json:"text"`
	LengthUnits    int             `json:"lengthUnits"`
	StartRuneIndex int             `json:"startRuneIndex"`
	EndRuneIndex   int             `json:"endRuneIndex"`
	StartByteIndex int             `json:"startByteIndex"`
	EndByteIndex   int             `json:"endByteIndex"`
	Characters     []CharacterInfo `json:"characters,omitempty"`
	UDH            *UDHInfo        `json:"udh,omitempty"`
}

type Segmentation struct {
	Encoding           Encoding     `json:"encoding"`
	LengthUnits        int          `json:"lengthUnits"`
	SegmentCount       int          `json:"segmentCount"`
	IsMultipart        bool         `json:"isMultipart"`
	SingleSegmentLimit int          `json:"singleSegmentLimit"`
	MultipartLimit     int          `json:"multipartLimit"`
	PerSegmentLimit    int          `json:"perSegmentLimit"`
	Segments           []SMSSegment `json:"segments,omitempty"`
	UDH                *UDHInfo     `json:"udh,omitempty"`
}

type UDHInfo struct {
	Required            bool   `json:"required"`
	HeaderLength        int    `json:"headerLength"`
	HeaderLengthSeptets int    `json:"headerLengthSeptets"`
	ReferenceNumber     int    `json:"referenceNumber"`
	ReferenceSize       int    `json:"referenceSize"`
	PartNumber          int    `json:"partNumber"`
	TotalParts          int    `json:"totalParts"`
	Hex                 string `json:"hex"`
}

type Warning struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	IndexRune  int    `json:"indexRune,omitempty"`
	Character  string `json:"character,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

type Suggestion struct {
	IndexRune   int    `json:"indexRune"`
	Original    string `json:"original"`
	Replacement string `json:"replacement"`
	Reason      string `json:"reason"`
	SavesUnits  int    `json:"savesUnits"`
}

type BillingInfo struct {
	Encoding      Encoding `json:"encoding"`
	SegmentCount  int      `json:"segmentCount"`
	BillableParts int      `json:"billableParts"`
	Reason        string   `json:"reason"`
}
