package array

import (
	"github.com/eandr-67/validator"
)

// Build реализует построитель валидатора массива
type Build struct {
	before, after validator.Rules[[]any]
	cell          validator.Builder
}

// Validator создаёт валидатор массива
func (b *Build) Validator() validator.Validator {
	return validator.NewFull[[]any](validator.Convert[[]any], b.before, b.after, &handle{b.cell.Validator()})
}

// Before добавляет действия в набор начальных действий
func (b *Build) Before(actions ...validator.Action[[]any]) *Build {
	b.before.Append(actions...)
	return b
}

// After добавляет действия в набор конечных действий
func (b *Build) After(actions ...validator.Action[[]any]) *Build {
	b.after.Append(actions...)
	return b
}
