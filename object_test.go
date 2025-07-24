package validator

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/stretchr/testify/assert"
)

func TestObj_Empty(t *testing.T) {
	vl := Obj().Compile()
	assert.Len(t, vl.(validator[map[string]any]).actions, 0)

	out, err := vl.Do(nil)
	assert.Nil(t, out)
	assert.Nil(t, err)

	out, err = vl.Do(map[string]any{"aaa": 45, "bbb": nil})
	assert.Equal(t, out, map[string]any{"aaa": 45, "bbb": nil})
	assert.Nil(t, err)

	var tmp map[string]any
	out, err = vl.Do(tmp)
	assert.Equal(t, out, map[string]any{})
	assert.Nil(t, err)

	out, err = vl.Do("aaa")
	assert.Equal(t, out, map[string]any{})
	assert.Equal(t, err, errs.Errors{"": {"type"}})
}

func TestObj_Panic(t *testing.T) {
	assert.PanicsWithValue(
		t, "field already exists", func() {
			Obj().Field("a", Int()).Field("a", Int())
		})

	assert.PanicsWithValue(
		t, "field already exists", func() {
			Obj().FieldList(map[string]Builder{"a": Int(), "c": Int()}).
				FieldList(map[string]Builder{"b": Int(), "c": Int()})
		})

	assert.PanicsWithValue(
		t, "default already exists", func() {
			Obj().Default("a", 25).Default("a", 25)
		})

	assert.PanicsWithValue(
		t, "default already exists", func() {
			Obj().DefaultList(map[string]any{"a": 25, "c": 25}).DefaultList(map[string]any{"b": 25, "c": 25})
		})

	assert.PanicsWithValue(
		t, "required field does unknown", func() {
			Obj().FieldList(map[string]Builder{"a": Int(), "c": Int(), "e": Int()}).
				Required("c", "d").Compile()
		})

	assert.PanicsWithValue(
		t, "both required and default", func() {
			Obj().FieldList(map[string]Builder{"a": Int(), "c": Int(), "e": Int()}).
				DefaultList(map[string]any{"a": Int(), "c": Int()}).
				Required("c", "e").Compile()
		})

	assert.PanicsWithValue(
		t, "default field does unknown", func() {
			Obj().FieldList(map[string]Builder{"a": Int(), "c": Int(), "e": Int()}).
				DefaultList(map[string]any{"b": Int(), "c": Int()}).Compile()
		})
}

func TestObj(t *testing.T) {
	out, err := Obj().Finish(NotNull).Compile().Do(nil)
	assert.Nil(t, out)
	assert.Equal(t, err, errs.Errors{"": {"null"}})

	out, err = Obj().Field("a", Int()).Finish(NotNull).Compile().Do(nil)
	assert.Nil(t, out)
	assert.Regexp(t, "^panic\\[", err[""][0])

	out, err = Obj().Field("a", Int()).Start(NotNull).Compile().Do(nil)
	assert.Nil(t, out)
	assert.Equal(t, err, errs.Errors{"": {"null"}})

	v := Obj(Null).Required("a").Default("b", int64(13)).
		FieldList(map[string]Builder{"a": String(), "b": Int(), "c": nil}).Compile()

	out, err = v.Do(nil)
	assert.Nil(t, out)
	assert.Nil(t, err)

	out, err = v.Do(map[string]any{})
	assert.Equal(t, out, map[string]any{"b": int64(13)})
	assert.Equal(t, err, errs.Errors{"a": {"missed"}})

	out, err = v.Do(map[string]any{"a": "x", "d": "y"})
	assert.Equal(t, out, map[string]any{"a": "x", "b": int64(13), "d": "y"})
	assert.Equal(t, err, errs.Errors{"d": {"unknown"}})

	out, err = v.Do(map[string]any{"a": "x", "b": "z"})
	assert.Equal(t, out, map[string]any{"a": "x", "b": int64(0)})
	assert.Equal(t, err, errs.Errors{"b": {"type"}})

	out, err = v.Do(map[string]any{"a": "x", "c": []any{1, "b", false}})
	assert.Equal(t, out, map[string]any{"a": "x", "b": int64(13), "c": []any{1, "b", false}})
	assert.Nil(t, err)
}
