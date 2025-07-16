package object

import "github.com/eandr-67/validator"

func Obj(rules ...validator.Action[map[string]any]) *Build {
	return &Build{
		before: rules,
		after:  validator.Rules[map[string]any]{},
		fields: map[string]validator.Builder{},
	}
}
