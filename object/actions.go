package object

import (
	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

func Default(field string, value any) validator.Action[map[string]any] {
	return func(val *map[string]any, _ string, _ *errs.Errors) (*map[string]any, bool) {
		if _, ok := (*val)[field]; !ok {
			(*val)[field] = value
		}
		return val, true
	}
}

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
