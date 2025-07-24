package validator

import (
	"github.com/eandr-67/errs"
)

// simpleConverter реализует простой преобразователь, используемый в тех случаях,
// когда any содержит либо значение целевого типа, либо nil и достаточно лишь конкретизировать тип any.
func simpleConverter[T any](raw any, err *errs.Errors) (value *T) {
	switch v := raw.(type) {
	case nil:
		return nil
	case T:
		return &v
	default:
		err.Add("", ErrMsg[ErrTypeIncorrect])
		var t T
		return &t
	}
}

// intConverter реализует преобразователь any -> int64, адаптированный для JSON, в котором все числа
// имеют вещественный тип и необходимо производить преобразование в 2 этапа: any -> float64 -> int64.
func intConverter(raw any, err *errs.Errors) (value *int64) {
	switch v := raw.(type) {
	case nil:
		return nil
	case int64:
		return &v
	case float64:
		t := int64(v)
		return &t
	default:
		err.Add("", ErrMsg[ErrTypeIncorrect])
		var t int64
		return &t
	}
}

// mapConverter реализует преобразователь any -> map[string]any.
//
// Map в Go хорош, когда уже инициализирован. А до этого момента он "немного" геморрой.
func mapConverter(raw any, err *errs.Errors) (value *map[string]any) {
	switch v := raw.(type) {
	case nil:
		return nil
	case map[string]any:
		if v != nil {
			return &v
		}
		return &map[string]any{}
	default:
		err.Add("", ErrMsg[ErrTypeIncorrect])
		return &map[string]any{}
	}
}
