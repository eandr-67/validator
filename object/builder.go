package object

import (
	"errors"

	"github.com/eandr-67/validator"
)

type Build struct {
	before, after validator.Rules[map[string]any]
	fields        map[string]validator.Builder
}

func (b *Build) Validator() validator.Validator {
	if len(b.fields) == 0 {
		panic(errors.New("fields is empty"))
	}
	h := make(handle, len(b.fields))
	for key, val := range b.fields {
		h[key] = val.Validator()
	}
	return validator.NewFull[map[string]any](validator.Convert[map[string]any], b.before, b.after, &h)
}

func (b *Build) Add(field string, build validator.Builder) *Build {
	if _, ok := b.fields[field]; ok {
		panic(errors.New("field is duplicated"))
	}
	b.fields[field] = build
	return b
}

func (b *Build) AddMap(fields map[string]validator.Builder) *Build {
	for key, val := range fields {
		b.Add(key, val)
	}
	return b
}

func (b *Build) Before(actions ...validator.Action[map[string]any]) *Build {
	b.before.Append(actions...)
	return b
}

func (b *Build) After(actions ...validator.Action[map[string]any]) *Build {
	b.after.Append(actions...)
	return b
}
