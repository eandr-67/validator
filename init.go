package validator

const (
	CodeTypeIncorrect = iota
	CodeFormatIncorrect
	CodeLengthIncorrect
	CodeValueIncorrect
	CodeValueIsNull
	CodeKeyMissed
	CodeKeyUnknown
)

var ErrMsg = []string{
	"type",
	"format",
	"length",
	"value",
	"is_null",
	"missed",
	"unknown",
}
