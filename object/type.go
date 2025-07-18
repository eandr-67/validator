package object

import "github.com/eandr-67/validator"

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
