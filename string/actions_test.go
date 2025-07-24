package string

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
	"github.com/stretchr/testify/assert"
)

func TestAction_Length(t *testing.T) {
	var err errs.Errors
	var str = []string{"ц", "цч", "цчш"}
	var cmp = map[string]struct {
		op  validator.Action[string]
		res [3]bool
	}{
		"LenEq":    {LenEq(2), [3]bool{false, true, false}},
		"LenNe":    {LenNe(2), [3]bool{true, false, true}},
		"LenLe":    {LenLe(2), [3]bool{true, true, false}},
		"LenGe":    {LenGe(2), [3]bool{false, true, true}},
		"LenIn":    {LenIn(1, 2), [3]bool{true, true, false}},
		"LenNotIn": {LenNotIn(1, 2), [3]bool{false, false, true}},
	}
	for k, v := range cmp {
		for i, p := range v.res {
			err = nil
			var inp = str[i]
			out, ok := v.op(&inp, &err)
			if p {
				assert.Truef(t, ok, k)
				assert.Equalf(t, *out, inp, k)
				assert.Nilf(t, err, k)
			} else {
				assert.Falsef(t, ok, k)
				assert.Equalf(t, *out, inp, k)
				assert.Equalf(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrLengthIncorrect]}}, k)
			}
		}
	}
}

func TestAction_Regex(t *testing.T) {
	var err errs.Errors
	var str = []string{"п5", "пп55", "ппп555", "пппп5555"}
	var cmp = map[string]struct {
		op  validator.Action[string]
		res [4]bool
	}{
		"Regex":    {Regex("^[а-я]{1,2}\\d{1,2}$"), [4]bool{true, true, false, false}},
		"NotRegex": {NotRegex("^[а-я]{2,3}\\d{2,3}$"), [4]bool{true, false, false, true}},
	}
	for k, v := range cmp {
		for i, p := range v.res {
			err = nil
			var inp = str[i]
			out, ok := v.op(&inp, &err)
			if p {
				assert.Truef(t, ok, k)
				assert.Equalf(t, *out, inp, k)
				assert.Nilf(t, err, k)
			} else {
				assert.Falsef(t, ok, k)
				assert.Equalf(t, *out, inp, k)
				assert.Equalf(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrFormatIncorrect]}}, k)
			}
		}
	}
}
