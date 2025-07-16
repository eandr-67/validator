package array

import (
	"errors"

	"github.com/eandr-67/validator"
)

func Arr(cell validator.Builder, rules ...validator.Action[[]any]) *Build {
	if cell == nil {
		panic(errors.New("cell cannot be nil"))
	}
	return &Build{
		before: rules,
		after:  validator.Rules[[]any]{},
		cell:   cell,
	}
}
