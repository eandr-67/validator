package array

import (
	"testing"

	v "github.com/eandr-67/validator"
)

func TestType(t *testing.T) {
	c := v.Int(v.NotNull, v.Gt(int64(10)), v.Lt(int64(20)))

	b := Arr(c).Before().After()
	if b.cell == nil {
		t.Errorf("cell = %v, want not nil", b.cell)
	}
	if b.before == nil {
		t.Errorf("before = %v, want not nil", b.before)
	}
	if b.after == nil {
		t.Errorf("after = %v, want not nil", b.after)
	}
}
