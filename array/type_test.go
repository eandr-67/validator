package array

import (
	"testing"

	v "github.com/eandr-67/validator"
)

func TestArr(t *testing.T) {
	c := v.Int(v.NotNull, v.Gt(int64(10)), v.Lt(int64(20)))

	b := Arr(c).Before().After()
	if b.cell == nil {
		t.Errorf("cell = %v, want not nil", b.cell)
	}
	if b.before == nil {
		t.Errorf("before = %v, want not nil", b.before)
	} else if len(b.before) != 0 {
		t.Errorf("before = %v, want empty", b.before)
	}
	if b.after == nil {
		t.Errorf("after = %v, want not nil", b.after)
	} else if len(b.after) != 0 {
		t.Errorf("after = %v, want empty", b.after)
	}

	b = Arr(c).Before(LenGe(2)).After(LenLe(3))
	if b.before == nil {
		t.Errorf("before = %v, want not nil", b.before)
	} else if len(b.before) != 1 {
		t.Errorf("before = %v, want len = 1", b.before)
	}
	if b.after == nil {
		t.Errorf("after = %v, want not nil", b.after)
	} else if len(b.after) != 1 {
		t.Errorf("after = %v, want len = 1", b.after)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("there should be panic")
		} else if v, ok := r.(error); !ok {
			t.Errorf("there should be error")
		} else if v.Error() != "cell cannot be nil" {
			t.Errorf("there should be 'cell cannot be nil'")
		}
	}()
	b = Arr(nil)
}
