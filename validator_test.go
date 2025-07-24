package validator

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/stretchr/testify/assert"
)

func TestNewValidator_Panic(t *testing.T) {
	assert.PanicsWithValue(
		t, "converter cannot be nil", func() {
			_ = NewValidator[any](nil)
		})
}

func TestNewValidator_Nil_Action(t *testing.T) {
	v := NewValidator(intConverter, nil, Null, nil, nil, Ge[int64](25))
	assert.Len(t, v.(validator[int64]).actions, 2)
}

func TestNewValidator_Nil_Data(t *testing.T) {
	v := NewValidator(intConverter, nil, Null, nil, nil, Ge[int64](25))

	out, err := v.Do(nil)
	assert.Nil(t, err)
	assert.Nil(t, out)

	out, err = v.Do(int64(125))
	assert.Nil(t, err)
	assert.Equal(t, out.(int64), int64(125))

	out, err = v.Do(float64(125))
	assert.Nil(t, err)
	assert.Equal(t, out.(int64), int64(125))

	out, err = v.Do(int(125))
	assert.Equal(t, err, errs.Errors{"": {"type"}})
	assert.Equal(t, out.(int64), int64(0))
}
