package validator

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/stretchr/testify/assert"
)

func TestAction_Null(t *testing.T) {
	var err errs.Errors

	out, ok := Null[int64](nil, &err)
	assert.False(t, ok)
	assert.Nil(t, out)
	assert.Nil(t, err)

	err = nil
	var inp int64 = 23
	out, ok = Null(&inp, &err)
	assert.True(t, ok)
	assert.Equal(t, *out, inp)
	assert.Nil(t, err)
}

func TestAction_IfNull(t *testing.T) {
	var err errs.Errors
	act := IfNull([]int64{2, 3})
	out, ok := act(&[]int64{1, 4}, &err)
	assert.True(t, ok)
	assert.Equal(t, *out, []int64{1, 4})
	assert.Nil(t, err)

	err = nil
	out, ok = act(nil, &err)
	assert.True(t, ok)
	assert.Equal(t, *out, []int64{2, 3})
	assert.Nil(t, err)
}

func TestAction_NotNull(t *testing.T) {
	var err errs.Errors
	var inp = "abc"

	out, ok := NotNull(&inp, &err)
	assert.True(t, ok)
	assert.Equal(t, *out, "abc")
	assert.Nil(t, err)

	err = nil
	out, ok = NotNull[string](nil, &err)
	assert.False(t, ok)
	assert.Nil(t, out)
	assert.Equal(t, err, errs.Errors{"": {ErrMsg[ErrValueIsNull]}})
}

func TestAction_Compare(t *testing.T) {
	var err = errs.Errors{}
	var cmp = map[string]struct {
		op  Action[int64]
		res [3]bool
	}{
		"Eq":    {Eq(int64(15)), [3]bool{false, true, false}},
		"Ne":    {Ne(int64(15)), [3]bool{true, false, true}},
		"Lt":    {Lt(int64(15)), [3]bool{true, false, false}},
		"Le":    {Le(int64(15)), [3]bool{true, true, false}},
		"Gt":    {Gt(int64(15)), [3]bool{false, false, true}},
		"Ge":    {Ge(int64(15)), [3]bool{false, true, true}},
		"In":    {In(int64(15), int64(16)), [3]bool{false, true, true}},
		"NotIn": {NotIn(int64(15), int64(16)), [3]bool{true, false, false}},
	}
	for k, v := range cmp {
		for i, p := range v.res {
			err = nil
			var inp = int64(i + 14)
			out, ok := v.op(&inp, &err)
			if p {
				assert.Truef(t, ok, k)
				assert.Equalf(t, *out, inp, k)
				assert.Nilf(t, err, k)
			} else {
				assert.Falsef(t, ok, k)
				assert.Equalf(t, *out, inp, k)
				assert.Equalf(t, err, errs.Errors{"": {ErrMsg[ErrValueIncorrect]}}, k)
			}
		}
	}
}
