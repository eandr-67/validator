package time

import (
	"slices"
	"testing"
	"time"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

func TestAction_Compare(t *testing.T) {
	var val [3]time.Time
	val[0], _ = time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:15", timeZone)
	val[1], _ = time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:16", timeZone)
	val[2], _ = time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:17", timeZone)
	var err = errs.Errors{}
	var cmp = map[string]struct {
		op  validator.Action[time.Time]
		res [3]bool
	}{
		"Eq":    {Eq(val[1]), [3]bool{false, true, false}},
		"Ne":    {Ne(val[1]), [3]bool{true, false, true}},
		"Lt":    {Lt(val[1]), [3]bool{true, false, false}},
		"Le":    {Le(val[1]), [3]bool{true, true, false}},
		"Gt":    {Gt(val[1]), [3]bool{false, false, true}},
		"Ge":    {Ge(val[1]), [3]bool{false, true, true}},
		"In":    {In(val[1], val[2]), [3]bool{false, true, true}},
		"NotIn": {NotIn(val[1], val[2]), [3]bool{true, false, false}},
	}
	for k, v := range cmp {
		for i, p := range v.res {
			err = errs.Errors{}
			var inp = val[i]
			out, ok := v.op(&inp, k, &err)
			if p {
				if len(err) != 0 {
					t.Errorf("%s(%v): len(err) should be 0", k, inp)
				}
				if *out != inp {
					t.Errorf("%s(%v): out must be %v, got %v", k, inp, inp, *out)
				}
				if !ok {
					t.Errorf("%s(%v): ok should be true", k, inp)
				}
			} else {
				checkError(t, err, k, []string{validator.ErrMsg[validator.CodeValueIncorrect]})
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
