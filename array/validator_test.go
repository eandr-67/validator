package array

import (
	"slices"
	"testing"

	"github.com/eandr-67/validator"
)

func TestBuild_Validator(t *testing.T) {
	c := validator.Int(validator.NotNull, validator.Gt(int64(10)), validator.Lt(int64(20)))
	b := Arr(c).Before(validator.Null, LenGe(2)).After(LenLe(3)).Validator()

	out, err := b.Do(nil)
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
	} else if len(*err) != 0 {
		t.Errorf("len(err) = %v, want 0", len(*err))
	}
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}

	out, err = b.Do(125)
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
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

	out, err = b.Do([]any{})
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
	} else if len(*err) != 1 {
		t.Errorf("len(err) = %v, want 1", len(*err))
	} else if e, ok := (*err)[""]; !ok {
		t.Errorf(`err = %v, want map[""]`, e)
	} else if len(e) != 1 {
		t.Errorf("len(e) = %v, want 1", len(e))
	} else if e[0] != "length" {
		t.Errorf("e[0] = %v, want length", e[0])
	}
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}

	out, err = b.Do([]any{int64(11), int64(12), int64(13), int64(14), int64(15)})
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
	} else if len(*err) != 1 {
		t.Errorf("len(err) = %v, want 1", len(*err))
	} else if e, ok := (*err)[""]; !ok {
		t.Errorf(`err = %v, want map[""]`, e)
	} else if len(e) != 1 {
		t.Errorf("len(e) = %v, want 1", len(e))
	} else if e[0] != "length" {
		t.Errorf("e[0] = %v, want length", e[0])
	}
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}

	out, err = b.Do([]any{int64(14), nil, int64(16)})
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
	} else if len(*err) != 1 {
		t.Errorf("len(err) = %v, want 1", len(*err))
	} else if e, ok := (*err)["[1]"]; !ok {
		t.Errorf(`err = %v, want map["[1]"]`, e)
	} else if len(e) != 1 {
		t.Errorf("len(e) = %v, want 1", len(e))
	} else if e[0] != "is_null" {
		t.Errorf("e[0] = %v, want is_null", e[0])
	}
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if !slices.Equal(out.([]any), []any{int64(14), nil, int64(16)}) {
		t.Errorf("out = %v, want []any{int64(14), nil, int64(16)}", out)
	}

	out, err = b.Do([]any{int64(14), int64(15), "16"})
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
	} else if len(*err) != 1 {
		t.Errorf("len(err) = %v, want 1", len(*err))
	} else if e, ok := (*err)["[2]"]; !ok {
		t.Errorf(`err = %v, want map["[2]"]`, e)
	} else if len(e) != 1 {
		t.Errorf("len(e) = %v, want 1", len(e))
	} else if e[0] != "type" {
		t.Errorf("e[0] = %v, want type", e[0])
	}
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if !slices.Equal(out.([]any), []any{int64(14), int64(15), nil}) {
		t.Errorf("out = %v, want []any{int64(14), int64(15), nil}", out)
	}

	out, err = b.Do([]any{float64(14), int64(15), float64(16)})
	if err == nil {
		t.Errorf("err = %v, want map[]", err)
	} else if len(*err) != 0 {
		t.Errorf("len(err) = %v, want 0", len(*err))
	}
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if v, ok := out.([]any); !ok {
		t.Errorf(`out = %v, want map[[]any]`, out)
	} else if !slices.Equal(v, []any{int64(14), int64(15), int64(16)}) {
		t.Errorf("out = %v, want []any{int64(14), int64(15), int64(16)}", out)
	}
}
