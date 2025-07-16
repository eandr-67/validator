package validator

import (
	"testing"

	"github.com/eandr-67/errs"
)

func TestRules_Append(t *testing.T) {
	var r Rules[int64]
	r.Append()
	if len(r) != 0 {
		t.Error("r should be empty")
	}
	r.Append(Null[int64], Gt(int64(10)))
	if len(r) != 2 {
		t.Error("len(r) should be 2")
	}
	r.Append(Lt(int64(15)))
	if len(r) != 3 {
		t.Error("len(r) should be 3")
	}
}

func TestRules_Execute(t *testing.T) {
	var err = errs.Errors{}
	var r Rules[int64]
	r.Append(Null[int64], Gt(int64(14)), Lt(int64(16)))

	out := r.Execute(nil, "", &err)
	if len(err) != 0 {
		t.Error("len(err) should be 0")
	}
	if out != nil {
		t.Error("out should be nil")
	}

	for k, v := range map[int64]bool{14: false, 15: true, 16: false} {
		err = errs.Errors{}
		out = r.Execute(&k, "Execute", &err)
		if v {
			if len(err) != 0 {
				t.Errorf("Execute(%d): len(err) should be 0", k)
			}
			if *out != k {
				t.Errorf("Execute(%d): out must be %d, got %d", k, k, *out)
			}
		} else {
			checkError(t, err, "Execute", []string{ErrMsg[CodeValueIncorrect]})
		}
	}
}
