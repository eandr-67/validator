package object

import (
	"encoding/json"
	"io"

	"github.com/eandr-67/errs"
	v "github.com/eandr-67/validator"
)

// Parse декодирует поток, содержащий JSON, в map[string]any и обрабатывает этот any валидатором.
// Возвращает результат обработки и список ошибок.
// Синтаксический сахар для того, чтобы не загромождать код API.
func Parse(reader io.Reader, validator v.Validator) (result map[string]any, err *errs.Errors) {
	var raw any
	if json.NewDecoder(reader).Decode(&raw) != nil {
		err = &errs.Errors{}
		err.Add("", v.ErrMsg[v.CodeFormatIncorrect])
		return nil, err
	}
	return anyToMap(validator.Do(raw))
}

// ParseString декодирует строку, содержащую JSON, в map[string]any и обрабатывает этот any валидатором.
// Возвращает результат обработки и список ошибок.
// Синтаксический сахар для того, чтобы не загромождать код API.
func ParseString(str string, validator v.Validator) (result any, err *errs.Errors) {
	var data any
	if json.Unmarshal([]byte(str), &data) != nil {
		err = &errs.Errors{}
		err.Add("", v.ErrMsg[v.CodeFormatIncorrect])
		return nil, err
	}
	return validator.Do(data)
}

// anyToMap конкретизирует any в map[string]any.
// При невозможности преобразования добавляет ошибку в список ошибок
func anyToMap(data any, err *errs.Errors) (map[string]any, *errs.Errors) {
	if res, ok := data.(map[string]any); ok {
		return res, err
	}
	err.Add("", v.ErrMsg[v.CodeTypeIncorrect])
	return nil, err
}
