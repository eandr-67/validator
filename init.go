package validator

const (
	ErrTypeIncorrect = iota
	ErrFormatIncorrect
	ErrLengthIncorrect
	ErrValueIncorrect
	ErrValueIsNull
	ErrKeyMissed
	ErrKeyUnknown
	ErrPanic
)

// ErrMsg содержит набор кодов ошибок. В любой момент коды (тексты) ошибок можно заменить на собственные.
var ErrMsg = []string{
	"type",
	"format",
	"length",
	"value",
	"null",
	"missed",
	"unknown",
	"panic[%#v]",
}
