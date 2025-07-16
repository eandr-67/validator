package array

import (
	"github.com/eandr-67/validator"
)

type Build struct {
	before, after validator.Rules[[]any]
	cell          validator.Builder
}

func (b *Build) Validator() validator.Validator {
	return validator.NewFull[[]any](validator.Convert[[]any], b.before, b.after, &handle{b.cell.Validator()})
}

func (b *Build) Before(actions ...validator.Action[[]any]) *Build {
	b.before.Append(actions...)
	return b
}

func (b *Build) After(actions ...validator.Action[[]any]) *Build {
	b.after.Append(actions...)
	return b
}
