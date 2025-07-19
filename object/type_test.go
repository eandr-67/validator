package object

import (
	"maps"
	"testing"

	"github.com/eandr-67/validator"
)

func TestObj(t *testing.T) {
	c := validator.Int(validator.NotNull, validator.Gt(int64(10)), validator.Lt(int64(20)))
	d := validator.String(validator.Null)

	b := Obj().Before().After()
	if b.fields == nil {
		t.Errorf("fields = %v, want not nil", b.fields)
	} else if len(b.fields) != 0 {
		t.Errorf("len(fields) = %v, want 0", len(b.fields))
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

	b = Obj(Required("p", "q")).Before(Default("c", int64(40))).After(DefaultList(map[string]any{"d": "aaa", "e": nil}))

	if b.fields == nil {
		t.Errorf("before = %v, want not nil", b.before)
	} else if len(b.fields) != 0 {
		t.Errorf("len(fields) = %v, want 0", len(b.fields))
	}
	if b.before == nil {
		t.Errorf("before = %v, want not nil", b.before)
	} else if len(b.before) != 2 {
		t.Errorf("before = %v, want len = 2", b.before)
	}
	if b.after == nil {
		t.Errorf("after = %v, want not nil", b.after)
	} else if len(b.after) != 1 {
		t.Errorf("after = %v, want len = 1", b.after)
	}

	b.AddMap(map[string]validator.Builder{"c": c, "d": d})
	if b.fields == nil {
		t.Errorf("before = %v, want not nil", b.before)
	} else if len(b.fields) != 2 {
		t.Errorf("len(fields) = %v, want 0", len(b.fields))
	} else {
		if v, ok := b.fields["c"]; !ok {
			t.Errorf(`fields["c"] = %v, want exists`, v)
		} else if v == nil {
			t.Errorf(`fields["c"] = %v, want not nil`, v)
		} else if q, ok := v.(*validator.Build[int64]); !ok {
			t.Errorf(`fields["c"] = %v, want Build[int64]`, q)
		}
		if v, ok := b.fields["d"]; !ok {
			t.Errorf(`fields["d"] = %v, want exists`, v)
		} else if v == nil {
			t.Errorf(`fields["d"] = %v, want not nil`, v)
		} else if q, ok := v.(*validator.Build[string]); !ok {
			t.Errorf(`fields["d"] = %v, want Build[int64]`, q)
		}
	}
	if b.before == nil {
		t.Errorf("before = %v, want not nil", b.before)
	} else if len(b.before) != 2 {
		t.Errorf("before = %v, want len = 2", b.before)
	}
	if b.after == nil {
		t.Errorf("after = %v, want not nil", b.after)
	} else if len(b.after) != 1 {
		t.Errorf("after = %v, want len = 1", b.after)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic expected")
		} else if v, ok := r.(error); !ok {
			t.Errorf(`r.(error) = %v, want error`, v)
		} else if v == nil {
			t.Errorf(`r.(error) = %v, want not nil`, v)
		} else if v.Error() != "field is duplicated" {
			t.Errorf(`v.Error() = %v, want "field is duplicated"`, v)
		}
	}()
	b.Add("c", d)
}

func TestObj_Validator_Empty(t *testing.T) {
	b := Obj(Required("p", "q")).Before(Default("c", int64(40))).
		After(DefaultList(map[string]any{"d": "aaa", "e": nil}))
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic expected")
		} else if v, ok := r.(error); !ok {
			t.Errorf(`r.(error) = %v, want error`, v)
		} else if v == nil {
			t.Errorf(`r.(error) = %v, want not nil`, v)
		} else if v.Error() != "fields is empty" {
			t.Errorf(`v.Error() = %v, want "fields is empty"`, v)
		}
	}()
	_ = b.Validator()
}

func TestObj_Validator(t *testing.T) {
	c := validator.Int(validator.NotNull, validator.Gt(int64(10)), validator.Lt(int64(20)))
	d := validator.String(validator.Null)

	b := Obj(validator.Null, Required("d")).After(Default("c", int64(40))).
		After(DefaultList(map[string]any{"e": nil})).
		Add("c", c).Add("d", d).Validator()

	out, err := b.Do(nil)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if err == nil {
		t.Errorf("err expected")
	} else if len(*err) != 0 {
		t.Errorf("len(*err) = %v, want 0", len(*err))
	}

	out, err = b.Do(125)
	if out != nil {
		t.Errorf("out = %v, want nil", out)
	}
	if err == nil {
		t.Errorf("err expected")
	} else if len(*err) != 1 {
		t.Errorf("len(*err) = %v, want 1", len(*err))
	} else if v, ok := (*err)[""]; !ok {
		t.Errorf(`(*err)[""] = %v, want exists`, v)
	} else if len(v) != 1 {
		t.Errorf(`(*err)[""] = %v, want len = 1`, v)
	} else if v[0] != "type" {
		t.Errorf(`(*err)[""] = %v, want "type"`, v)
	}

	out, err = b.Do(map[string]any{})
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if v, ok := out.(map[string]any); !ok {
		t.Errorf(`out.(map[string]any) = %v, want exists`, v)
	} else if !maps.Equal(v, map[string]any{"c": int64(40), "e": nil}) {
		t.Errorf(`v = %v, want %v`, v, map[string]any{"c": int64(40), "e": nil})
	}
	if err == nil {
		t.Errorf("err expected")
	} else if len(*err) != 1 {
		t.Errorf("len(*err) = %v, want 1", len(*err))
	} else if v, ok := (*err)[".d"]; !ok {
		t.Errorf(`(*err)[".d"] = %v, want exists`, v)
	} else if len(v) != 1 {
		t.Errorf(`(*err)[".d"] = %v, want len = 1`, v)
	} else if v[0] != "missed" {
		t.Errorf(`(*err)[""] = %v, want "missed"`, v)
	}

	out, err = b.Do(map[string]any{"d": nil, "c": int64(3)})
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if v, ok := out.(map[string]any); !ok {
		t.Errorf(`out.(map[string]any) = %v, want exists`, v)
	} else if !maps.Equal(v, map[string]any{"c": int64(3), "d": nil, "e": nil}) {
		t.Errorf(`v = %v, want %v`, v, map[string]any{"c": int64(3), "d": nil, "e": nil})
	}
	if err == nil {
		t.Errorf("err expected")
	} else if len(*err) != 1 {
		t.Errorf("len(*err) = %v, want 1", len(*err))
	} else if v, ok := (*err)[".c"]; !ok {
		t.Errorf(`(*err)[".c"] = %v, want exists`, v)
	} else if len(v) != 1 {
		t.Errorf(`(*err)[".c"] = %v, want len = 1`, v)
	} else if v[0] != "value" {
		t.Errorf(`(*err)[""] = %v, want "value"`, v)
	}

	out, err = b.Do(map[string]any{"d": nil, "q": 17})
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if v, ok := out.(map[string]any); !ok {
		t.Errorf(`out.(map[string]any) = %v, want exists`, v)
	} else if !maps.Equal(v, map[string]any{"c": int64(40), "d": nil, "e": nil, "q": 17}) {
		t.Errorf(`v = %v, want %v`, v, map[string]any{"c": int64(40), "d": nil, "e": nil, "q": 17})
	}
	if err == nil {
		t.Errorf("err expected")
	} else if len(*err) != 1 {
		t.Errorf("len(*err) = %v, want 1", len(*err))
	} else if v, ok := (*err)[".q"]; !ok {
		t.Errorf(`(*err)[".q"] = %v, want exists`, v)
	} else if len(v) != 1 {
		t.Errorf(`(*err)[".q"] = %v, want len = 1`, v)
	} else if v[0] != "unknown" {
		t.Errorf(`(*err)[""] = %v, want "unknown"`, v)
	}

	out, err = b.Do(map[string]any{"d": "aaa", "c": int64(15)})
	if out == nil {
		t.Errorf("out = %v, want not nil", out)
	} else if v, ok := out.(map[string]any); !ok {
		t.Errorf(`out.(map[string]any) = %v, want exists`, v)
	} else if !maps.Equal(v, map[string]any{"c": int64(15), "d": "aaa", "e": nil}) {
		t.Errorf(`v = %v, want %v`, v, map[string]any{"c": int64(15), "d": "aaa", "e": nil})
	}
	if err == nil {
		t.Errorf("err expected")
	} else if len(*err) != 0 {
		t.Errorf("len(*err) = %v, want 0", len(*err))
	}
}
