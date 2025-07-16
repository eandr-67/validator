package validator

import (
	"testing"
)

func TestSimple(t *testing.T) {
	v := Int(Null, Gt[int64](14), Lt(int64(16))).Validator()

	out, err := v.Do("123")
	if out != nil {
		t.Error("out should be nil")
	}
	checkError(t, *err, "", []string{ErrMsg[CodeTypeIncorrect]})

	out, err = v.Do(nil)
	if out != nil {
		t.Error("out should be nil")
	}
	if len(*err) != 0 {
		t.Error("len(err) should be 0")
	}

	for i := int64(14); i <= 16; i++ {
		out, err = v.Do(i)
		if out == nil {
			t.Error("out should not be nil")
		} else if out.(int64) != i {
			t.Error("out.(int64) should be equal to i")
		}
		if i != 15 {
			checkError(t, *err, "", []string{ErrMsg[CodeValueIncorrect]})
		} else if len(*err) != 0 {
			t.Error("len(err) should be 0")
		}
	}
}
