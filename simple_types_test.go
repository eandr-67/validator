package validator

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/stretchr/testify/assert"
)

func TestFactory_Ok(t *testing.T) {
	tst := map[string]struct {
		builder Builder
		value   any
	}{
		"Int":      {Int().Add(NotNull), int64(25)},
		"Float":    {Float().Add(NotNull), 25.0},
		"String":   {String().Add(NotNull), "25"},
		"Bool":     {Bool().Add(NotNull), true},
		"Any":      {Any().Add(NotNull), []any{"a", false}},
		"Any(nil)": {Any().Add(), nil},
	}

	for n, f := range tst {
		out, err := f.builder.Compile().Do(f.value)
		assert.Nilf(t, err, n)
		assert.Equalf(t, f.value, out, n)
	}
}

func TestFactory_Bad(t *testing.T) {
	tst := map[string]struct {
		builder Builder
		value   any
	}{
		"Int":    {Int().Add(NotNull), 25},
		"Float":  {Float().Add(NotNull), 25},
		"String": {String().Add(NotNull), 25},
		"Bool":   {Bool().Add(NotNull), 25},
	}

	for n, f := range tst {
		out, err := f.builder.Compile().Do(f.value)
		assert.Equalf(t, err, errs.Errors{"": {"type"}}, n)
		assert.Emptyf(t, out, n)
	}
}

func TestFactory_Nil(t *testing.T) {
	tst := map[string]Builder{
		"Int":    Int().Add(Null),
		"Float":  Float().Add(Null),
		"String": String().Add(Null),
		"Bool":   Bool().Add(Null),
		"Any":    Any().Add(Null),
	}

	for n, f := range tst {
		out, err := f.Compile().Do(nil)
		assert.Nilf(t, err, n)
		assert.Nilf(t, out, n)
	}
}
