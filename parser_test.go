package validator

import (
	"strings"
	"testing"

	"github.com/eandr-67/errs"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	vl := Int(Gt[int64](5)).Compile()

	assert.Panics(t, func() { Parse(nil, vl) })

	out, err := Parse(strings.NewReader(""), vl)
	assert.Equal(t, err, errs.Errors{"": {"format"}})
	assert.Nil(t, out)

	out, err = Parse(strings.NewReader("null"), vl)
	assert.Regexp(t, "^panic\\[", err[""][0])
	assert.Nil(t, out)

	out, err = Parse(strings.NewReader(`"aaa"`), vl)
	assert.Equal(t, err, errs.Errors{"": {"type"}})
	assert.Equal(t, out.(int64), int64(0))

	out, err = Parse(strings.NewReader("25"), vl)
	assert.Nil(t, err)
	assert.Equal(t, out.(int64), int64(25))

	out, err = Parse(strings.NewReader("25.7"), vl)
	assert.Nil(t, err)
	assert.Equal(t, out.(int64), int64(25))

	out, err = Parse(strings.NewReader("-25"), vl)
	assert.Equal(t, err, errs.Errors{"": {"value"}})
	assert.Equal(t, out.(int64), int64(-25))
}

func TestParseStr(t *testing.T) {
	vl := Int(Gt[int64](5)).Compile()

	out, err := ParseStr("", vl)
	assert.Equal(t, err, errs.Errors{"": {"format"}})
	assert.Nil(t, out)

	out, err = ParseStr("null", vl)
	assert.Regexp(t, "^panic\\[", err[""][0])
	assert.Nil(t, out)

	out, err = ParseStr(`"aaa"`, vl)
	assert.Equal(t, err, errs.Errors{"": {"type"}})
	assert.Equal(t, out.(int64), int64(0))

	out, err = ParseStr("25", vl)
	assert.Nil(t, err)
	assert.Equal(t, out.(int64), int64(25))

	out, err = ParseStr("25.7", vl)
	assert.Nil(t, err)
	assert.Equal(t, out.(int64), int64(25))

	out, err = ParseStr("-25", vl)
	assert.Equal(t, err, errs.Errors{"": {"value"}})
	assert.Equal(t, out.(int64), int64(-25))
}
