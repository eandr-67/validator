package uuid

import (
	"testing"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUuidConverter(t *testing.T) {
	var err errs.Errors
	var inp uuid.UUID

	out := uuidConverter(nil, &err)
	assert.Nil(t, err)
	assert.Nil(t, out)

	err = nil
	tmp, _ := uuid.Parse("12345678-9abc-def0-1234-56789abcdef0")
	out = uuidConverter("12345678-9abc-def0-1234-56789abcdef0", &err)
	assert.Nil(t, err)
	assert.Equal(t, *out, tmp)

	err = nil
	out = uuidConverter("123456789abcdef0123456789abcdef0", &err)
	assert.Nil(t, err)
	assert.Equal(t, *out, tmp)

	err = nil
	out = uuidConverter("urn:uuid:12345678-9abc-def0-1234-56789abcdef0", &err)
	assert.Nil(t, err)
	assert.Equal(t, *out, tmp)

	err = nil
	out = uuidConverter("{12345678-9abc-def0-1234-56789abcdef0}", &err)
	assert.Nil(t, err)
	assert.Equal(t, *out, tmp)

	err = nil
	out = uuidConverter("2025-07-01 23:15", &err)
	assert.Equal(t, *out, inp)
	assert.Equal(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrFormatIncorrect]}})

	err = nil
	out = uuidConverter(25, &err)
	assert.Equal(t, *out, inp)
	assert.Equal(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrTypeIncorrect]}})
}
