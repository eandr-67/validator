package object

import (
	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// Default добавляет в объект поле field со значением value, если в обрабатываемом объекте поля с таким именем нет.
func Default(field string, value any) validator.Action[map[string]any] {
	return func(val *map[string]any, _ string, _ *errs.Errors) (*map[string]any, bool) {
		if _, ok := (*val)[field]; !ok {
			(*val)[field] = value
		}
		return val, true
	}
}

// DefaultList добавляет в объект те поля с именами - ключами fields и значениями - значениями fields,
// которые в обрабатываемом объекте отсутствуют.
func DefaultList(fields map[string]any) validator.Action[map[string]any] {
	return func(val *map[string]any, _ string, _ *errs.Errors) (*map[string]any, bool) {
		for key, value := range fields {
			if _, ok := (*val)[key]; !ok {
				(*val)[key] = value
			}
		}
		return val, true
	}
}

// Required проверяет существование в объекте полей с именами fields.
// Для каждого отсутствующего поля генерируется ошибка с ключом - именем поля.
func Required(fields ...string) validator.Action[map[string]any] {
	return func(val *map[string]any, key string, err *errs.Errors) (*map[string]any, bool) {
		for _, field := range fields {
			if _, ok := (*val)[field]; !ok {
				err.Add(key+"."+field, validator.ErrMsg[validator.CodeKeyMissed])
			}
		}
		return val, true
	}
}
