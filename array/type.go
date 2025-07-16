package array

import (
	"errors"

	"github.com/eandr-67/validator"
)

func Arr(cell validator.Builder, before ...validator.Action[[]any]) *Build {
	if cell == nil {
		panic(errors.New("cell cannot be nil"))
	}
	return &Build{
		before: before,
		after:  validator.Rules[[]any]{},
		cell:   cell,
	}
}
