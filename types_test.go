package validator

import (
	"testing"
)

func TestInt(t *testing.T) {
	b := Int(Gt[int64](13), Lt(int64(25))).Append()
	if len(b.Rules) != 2 {
		t.Errorf("Rule length should be 2")
	}
}

func TestFloat(t *testing.T) {
	b := Float(Gt(13.7)).Append(Lt(25.2))
	if len(b.Rules) != 2 {
		t.Errorf("Rule length should be 2")
	}
}

func TestString(t *testing.T) {
	b := String().Append(Gt("13"), Lt("25"))
	if len(b.Rules) != 2 {
		t.Errorf("Rule length should be 2")
	}
}

func TestBool(t *testing.T) {
	b := Bool(Null, Eq(true))
	if len(b.Rules) != 2 {
		t.Errorf("Rule length should be 2")
	}
}

func TestEmpty(t *testing.T) {
	b := Bool()
	if b.Rules == nil {
		t.Errorf("Rules should not be nil")
	}
	if len(b.Rules) != 0 {
		t.Errorf("Rule length should be 0")
	}
}
