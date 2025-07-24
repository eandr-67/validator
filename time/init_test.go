package time

import (
	"testing"
	"time"

	"github.com/eandr-67/errs"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	var err errs.Errors
	ct := timeConverter([]string{"2006-01-02 15:04:05.999999999"})
	tz, _ := time.LoadLocation("Europe/Moscow")

	out1 := ct("2025-07-06 05:04:03", &err)
	assert.Nil(t, err)

	SetTimeZone(tz)
	out2 := ct("2025-07-06 08:04:03", &err)
	assert.Nil(t, err)
	assert.True(t, out1.Equal(*out2))
}
