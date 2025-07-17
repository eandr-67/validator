package validator

import (
	"testing"
)

func TestNewBuilder(t *testing.T) {
	b := NewBuilder(Convert[string], nil)
	if b.rules == nil {
		t.Error("rules cannot be nil")
	} else if len(b.rules) != 0 {
		t.Error("rules should be empty")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("there should be panic")
		} else if v, ok := r.(error); !ok {
			t.Errorf("there must be error")
		} else if v.Error() != "convert cannot be nil" {
			t.Errorf(`error message must be "convert cannot be nil", got "%s"`, v.Error())
		}
	}()

	b = NewBuilder[string](nil, nil)
}
