package validator

import (
	"slices"
	"testing"

	"github.com/eandr-67/errs"
)

func TestAction_Null(t *testing.T) {
	var err errs.Errors

	out, ok := Null[int64](nil, "", &err)
	if err != nil {
		t.Error("err should be nil")
	}
	if out != nil {
		t.Error("out should be nil")
	}
	if ok {
		t.Error("ok should be false")
	}

	err = nil
	var inp int64 = 23
	out, ok = Null(&inp, "", &err)
	if err != nil {
		t.Error("err should be nil")
	}
	if out == nil {
		t.Error("out should not be nil")
	} else if *out != inp {
		t.Error("out should be equal to inp")
	}
	if !ok {
		t.Error("ok should be true")
	}
}

func TestAction_IfNull(t *testing.T) {
	var err errs.Errors
	act := IfNull([]int64{2, 3})
	out, ok := act(&[]int64{1, 4}, "", &err)
	if err != nil {
		t.Error("err should be nil")
	}
	if out == nil {
		t.Error("out should not be nil")
	} else if !slices.Equal(*out, []int64{1, 4}) {
		t.Errorf("out must be %#v, got %#v", []int64{1, 4}, *out)
	}
	if !ok {
		t.Error("ok should be true")
	}

	err = nil
	out, ok = act(nil, "", &err)
	if err != nil {
		t.Error("err should be nil")
	}
	if out == nil {
		t.Error("out should not be nil")
	} else if !slices.Equal(*out, []int64{2, 3}) {
		t.Errorf("out must be %#v, got %#v", []int64{2, 3}, *out)
	}
	if !ok {
		t.Error("ok should be true")
	}
}

func TestAction_NotNull(t *testing.T) {
	var err = errs.Errors{}
	var inp = "abc"

	out, ok := NotNull(&inp, "@", &err)
	if len(err) != 0 {
		t.Error("len(err) should be 0")
	}
	if out == nil {
		t.Error("out should not be nil")
	} else if *out != "abc" {
		t.Error("out should be 'abc'")
	}
	if !ok {
		t.Error("ok should be true")
	}

	err = errs.Errors{}
	out, ok = NotNull[string](nil, "@", &err)
	checkError(t, err, "@", []string{ErrMsg[CodeValueIsNull]})
	if out != nil {
		t.Error("out should be nil")
	}
	if ok {
		t.Error("ok should be false")
	}
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
			err = errs.Errors{}
			var inp = int64(i + 14)
			out, ok := v.op(&inp, k, &err)
			if p {
				if len(err) != 0 {
					t.Errorf("%s(%d): len(err) should be 0", k, inp)
				}
				if *out != inp {
					t.Errorf("%s(%d): out must be %d, got %d", k, inp, inp, *out)
				}
				if !ok {
					t.Errorf("%s(%d): ok should be true", k, inp)
				}
			} else {
				checkError(t, err, k, []string{ErrMsg[CodeValueIncorrect]})
				if ok {
					t.Errorf("%s(%d): ok should be false", k, inp)
				}
			}
		}
	}
}

func TestAction_Length(t *testing.T) {
	var err = errs.Errors{}
	var str = []string{"a", "ab", "abc"}
	var cmp = map[string]struct {
		op  Action[string]
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
			var inp = str[i]
			out, ok := v.op(&inp, k, &err)
			if p {
				if len(err) != 0 {
					t.Errorf("%s(%s): err should be nil", k, inp)
				}
				if *out != inp {
					t.Errorf("%s(%s): out must be %s, got %s", k, inp, inp, *out)
				}
				if !ok {
					t.Errorf("%s(%s): ok should be true", k, inp)
				}
			} else {
				checkError(t, err, k, []string{ErrMsg[CodeLengthIncorrect]})
				if ok {
					t.Errorf("%s(%s): ok should be false", k, inp)
				}
			}
		}
	}
}

func TestAction_Regex(t *testing.T) {
	var err = errs.Errors{}
	var str = []string{"п5", "пп55", "ппп555", "пппп5555"}
	var cmp = map[string]struct {
		op  Action[string]
		res [4]bool
	}{
		"Regex":    {Regex("^[а-я]{1,2}\\d{1,2}$"), [4]bool{true, true, false, false}},
		"NotRegex": {NotRegex("^[а-я]{2,3}\\d{2,3}$"), [4]bool{true, false, false, true}},
	}
	for k, v := range cmp {
		for i, p := range v.res {
			err = errs.Errors{}
			var inp = str[i]
			out, ok := v.op(&inp, k, &err)
			if p {
				if len(err) != 0 {
					t.Errorf("%s(%s): err should be nil", k, inp)
				}
				if *out != inp {
					t.Errorf("%s(%s): out must be %s, got %s", k, inp, inp, *out)
				}
				if !ok {
					t.Errorf("%s(%s): ok should be true", k, inp)
				}
			} else {
				checkError(t, err, k, []string{ErrMsg[CodeFormatIncorrect]})
				if ok {
					t.Errorf("%s(%s): ok should be false", k, inp)
				}
			}
		}
	}
}

func TestAction_In_Empty(t *testing.T) {
	testIn(t, In[int], "In")
	testIn(t, NotIn[int], "NotIn")
	testIn(t, LenIn, "LenIn")
	testIn(t, LenNotIn, "LenNotIn")
}

func TestAction_Len_Negative(t *testing.T) {
	testLen(t, LenEq, "LenEq")
	testLen(t, LenNe, "LenNe")
	testLen(t, LenGe, "LenGe")
	testLen(t, LenLe, "LenLe")
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

func testIn[Q any](t *testing.T, f func(values ...int) Action[Q], name string) {
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

func testLen(t *testing.T, f func(value int) Action[string], name string) {
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
