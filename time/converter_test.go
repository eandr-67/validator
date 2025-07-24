package time

import (
	"testing"
	"time"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
	"github.com/stretchr/testify/assert"
)

func TestTimeConverter(t *testing.T) {
	var err errs.Errors
	cnv := timeConverter(Default)

	out := cnv(nil, &err)
	assert.Nil(t, err)
	assert.Nil(t, out)

	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:16", timeZone)
	out = cnv("2025-07-01 23:15:16+03:00", &err)
	assert.Nil(t, err)
	assert.True(t, tmp.Equal(*out))

	out = cnv("2025-07-01 23:15", &err)
	assert.Empty(t, *out)
	assert.Equal(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrFormatIncorrect]}})

	err = nil
	out = cnv(25, &err)
	assert.Empty(t, *out)
	assert.Equal(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrTypeIncorrect]}})
}
