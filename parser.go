package validator

import (
	"encoding/json"
	"io"

	"github.com/eandr-67/errs"
)

// Parse декодирует поток reader, содержащий JSON, any и обрабатывает этот any валидатором vl.
// Возвращает результат обработки и список ошибок.
func Parse(reader io.Reader, vl Validator) (result any, err errs.Errors) {
	var raw any

	if json.NewDecoder(reader).Decode(&raw) != nil {
		err.Add("", ErrMsg[ErrFormatIncorrect])
		return nil, err
	}
	return vl.Do(raw)
}

// ParseStr декодирует строку str, содержащую JSON, в any и обрабатывает этот any валидатором vl.
// Возвращает результат обработки и список ошибок.
func ParseStr(str string, vl Validator) (result any, err errs.Errors) {
	var raw any

	if json.Unmarshal([]byte(str), &raw) != nil {
		err.Add("", ErrMsg[ErrFormatIncorrect])
		return nil, err
	}
	return vl.Do(raw)
}
