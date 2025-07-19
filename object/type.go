package object

import "github.com/eandr-67/validator"

// Obj создаёт построитель валидатора массива. Получает действия, добавляемые в набор начальных действий.
func Obj(before ...validator.Action[map[string]any]) *Build {
	if before == nil {
		before = []validator.Action[map[string]any]{}
	}
	return &Build{
		before: before,
		after:  validator.Rules[map[string]any]{},
		fields: map[string]validator.Builder{},
	}
}
