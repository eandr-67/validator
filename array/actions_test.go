package array

import (
	"slices"
	"testing"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

func TestAction_Length(t *testing.T) {
	var err = errs.Errors{}
	var arr = [][]any{{"a"}, {"a", "ab"}, {"a", "ab", "abc"}}
	var cmp = map[string]struct {
		op  validator.Action[[]any]
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
			err = errs.Errors{}
			var inp = arr[i]
			out, ok := v.op(&inp, k, &err)
			if p {
				if len(err) != 0 {
					t.Errorf("%s(%v): err should be nil", k, inp)
				} else if !slices.Equal(*out, inp) {
					t.Errorf("%s(%v): out must be %v, got %v", k, inp, inp, *out)
				}
				if !ok {
					t.Errorf("%s(%v): ok should be true", k, inp)
				}
			} else {
				checkError(t, err, k, []string{validator.ErrMsg[validator.CodeLengthIncorrect]})
				if ok {
					t.Errorf("%s(%v): ok should be false", k, inp)
				}
			}
		}
	}
}

func checkError(t *testing.T, err errs.Errors, key string, val []string) {
	if err == nil {
		t.Error("err should not be nil")
	} else if len(err) != 1 {
		t.Error("err should have length 1")
	} else if v, ok := err[key]; !ok {
		t.Errorf("err should have key '%s'", key)
	} else if v == nil {
		t.Errorf("err[%s] should not be nil", key)
	} else if !slices.Equal(v, val) {
		t.Errorf("err[%s] must be %#v, got %#v", key, val, v)
	}
}

func TestAction_In_Empty(t *testing.T) {
	testIn(t, LenIn, "LenIn")
	testIn(t, LenNotIn, "LenNotIn")
}

func TestAction_Len_Negative(t *testing.T) {
	testLen(t, LenEq, "LenEq")
	testLen(t, LenNe, "LenNe")
	testLen(t, LenGe, "LenGe")
	testLen(t, LenLe, "LenLe")
}

func testIn(t *testing.T, f func(values ...int) validator.Action[[]any], name string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: there should be panic", name)
		} else if v, ok := r.(error); !ok {
			t.Errorf("%s: there must be error", name)
		} else if v.Error() != "must have at least one value" {
			t.Errorf(`%s: error message must be "must have at least one value", got "%s"`, name, v.Error())
		}
	}()
	_ = f()
}

func testLen(t *testing.T, f func(value int) validator.Action[[]any], name string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: there should be panic", name)
		} else if v, ok := r.(error); !ok {
			t.Errorf("%s: there must be error", name)
		} else if v.Error() != "the value must be not less than 0" {
			t.Errorf(`%s: error message must be "the value must be not less than 0", got "%s"`, name, v.Error())
		}
	}()
	_ = f(-1)
}
