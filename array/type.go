package array

import (
	"errors"

	"github.com/eandr-67/validator"
)

// Arr создаёт построитель валидатора массива. Получает построитель валидатора элементов массива cell и
// действия, добавляемые в набор начальных действий.
func Arr(cell validator.Builder, before ...validator.Action[[]any]) *Build {
	if cell == nil {
		panic(errors.New("cell cannot be nil"))
	}
	if before == nil {
		before = []validator.Action[[]any]{}
	}
	return &Build{
		before: before,
		after:  validator.Rules[[]any]{},
		cell:   cell,
	}
}
