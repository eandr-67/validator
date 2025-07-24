package array

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
	"github.com/stretchr/testify/assert"
)

func TestArr_Empty(t *testing.T) {
	v := validator.Arr(nil).Compile()

	out, err := v.Do(nil)
	assert.Nil(t, out)
	assert.Nil(t, err)

	out, err = v.Do([]any{"aaa", 45})
	assert.Equal(t, out, []any{"aaa", 45})
	assert.Nil(t, err)
}

func TestArr(t *testing.T) {
	v := validator.Arr(validator.String(validator.NotNull), validator.NotNull).
		Start(LenGe(2)).Finish(LenLe(3)).Compile()

	out, err := v.Do(nil)
	assert.Nil(t, out)
	assert.Equal(t, err, errs.Errors{"": {"null"}})

	out, err = v.Do([]any{"aaa", 45, "bbb"})
	assert.Equal(t, out, []any{"aaa", "", "bbb"})
	assert.Equal(t, err, errs.Errors{"1": {"type"}})

	out, err = v.Do([]any{45})
	assert.Equal(t, out, []any{45})
	assert.Equal(t, err, errs.Errors{"": {"length"}})

	out, err = v.Do([]any{"aaa", 45, "bbb", nil})
	assert.Equal(t, out, []any{"aaa", "", "bbb", nil})
	assert.Equal(t, err, errs.Errors{"": {"length"}, "1": {"type"}, "3": {"null"}})
}
