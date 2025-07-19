package object

import (
	"maps"
	"strings"
	"testing"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

func TestAnyToMap(t *testing.T) {
	err := &errs.Errors{}
	out, err := anyToMap(nil, err)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if err == nil {
		t.Errorf("err = <nil>")
	} else if v, ok := (*err)[""]; !ok {
		t.Errorf(`err[""] = %v, want not nil`, v)
	} else if len(v) != 1 {
		t.Errorf(`len(v) = %v, want 1`, len(v))
	} else if v[0] != "type" {
		t.Errorf(`v[0] = %v, want "type"`, v[0])
	}

	err = &errs.Errors{}
	out, err = anyToMap(map[string]any{}, err)
	if out == nil {
		t.Errorf("out = %v, want map[string]any{}", out)
	} else if len(out) != 0 {
		t.Errorf(`len(out) = %v, want 0`, len(out))
	}
	if err == nil {
		t.Errorf(`err = %v, want not nil`, err)
	} else if len(*err) != 0 {
		t.Errorf(`len(*err) = %v, want 0`, len(*err))
	}
}

func TestParseString(t *testing.T) {
	b := Obj(Required("a")).Add("a", validator.Int()).Validator()

	out, err := ParseString("", b)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if (*err)[""][0] != "format" {
		t.Errorf(`(*err)[""][0] = %v, want "format"`, (*err)[""][0])
	}

	out, err = ParseString("abc", b)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if (*err)[""][0] != "format" {
		t.Errorf(`(*err)[""][0] = %v, want "format"`, (*err)[""][0])
	}

	out, err = ParseString("{}", b)
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if len(out) != 0 {
		t.Errorf(`len(out) = %v, want 0`, len(out))
	}
	if (*err)[".a"][0] != "missed" {
		t.Errorf(`(*err)[".a"][0] = %v, want "missed"`, (*err)[".a"][0])
	}

	out, err = ParseString(`{"a": 25}`, b)
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if !maps.Equal(out, map[string]any{"a": int64(25)}) {
		t.Errorf(`out = %v, want map[string]any{"a": 25}`, out)
	}
	if len(*err) != 0 {
		t.Errorf(`len(*err) = %v, want nil`, len(*err))
	}
}

func TestParse(t *testing.T) {
	b := Obj(Required("a")).Add("a", validator.Int()).Validator()

	out, err := Parse(strings.NewReader(""), b)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if (*err)[""][0] != "format" {
		t.Errorf(`(*err)[""][0] = %v, want "format"`, (*err)[""][0])
	}

	out, err = Parse(strings.NewReader("abc"), b)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if (*err)[""][0] != "format" {
		t.Errorf(`(*err)[""][0] = %v, want "format"`, (*err)[""][0])
	}

	out, err = Parse(strings.NewReader("{}"), b)
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if len(out) != 0 {
		t.Errorf(`len(out) = %v, want 0`, len(out))
	}
	if (*err)[".a"][0] != "missed" {
		t.Errorf(`(*err)[".a"][0] = %v, want "missed"`, (*err)[".a"][0])
	}

	out, err = Parse(strings.NewReader(`{"a": 25}`), b)
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if !maps.Equal(out, map[string]any{"a": int64(25)}) {
		t.Errorf(`out = %v, want map[string]any{"a": 25}`, out)
	}
	if len(*err) != 0 {
		t.Errorf(`len(*err) = %v, want nil`, len(*err))
	}
}
