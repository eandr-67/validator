package array

import (
	"testing"

	"github.com/eandr-67/validator"
)

func TestBuild_Validator(t *testing.T) {
	c := validator.Int(validator.NotNull, validator.Gt(int64(10)), validator.Lt(int64(20)))
	b := Arr(c).Before(validator.Null, LenGe(2)).After(LenLe(3)).Validator()

	out, err := b.Do(nil)
	if err == nil {
		t.Errorf("err = %v, want map[]", *err)
	} else if len(*err) != 0 {
		t.Errorf("len(err) = %v, want 0", len(*err))
	}
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}

	out, err = b.Do(125)
	if err == nil {
		t.Errorf("err = %v, want map[]", *err)
	} else if len(*err) != 1 {
		t.Errorf("len(err) = %v, want 1", len(*err))
	} else if e, ok := (*err)[""]; !ok {
		t.Errorf(`err = %v, want map[""]`, e)
	} else if len(e) != 1 {
		t.Errorf("len(e) = %v, want 1", len(e))
	} else if e[0] != "type" {
		t.Errorf("e[0] = %v, want type", e[0])
	}
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}

	out, err = b.Do([]int{})
}
