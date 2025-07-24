package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime_OK(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:15", timeZone)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:17", timeZone)
	b := Time([]string{"2006-01-02 15:04:05"}, Gt(t1)).Add(Lt(t2))

	assert.Len(t, b.actions, 2)
	assert.Len(t, b.formats, 1)
	assert.NotPanics(t, func() { _ = b.Compile() })

}

func TestTime_NoFormats(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:15", timeZone)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:17", timeZone)
	assert.PanicsWithValue(t, "formats cannot be empty", func() { _ = Time([]string{}, Gt(t1)).Add(Lt(t2)).Compile() })
}
