package time

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	ct := convertTime([]string{"2006-01-02 15:04:05.999999999"})
	tz, _ := time.LoadLocation("Europe/Moscow")

	out1, err := ct("2025-07-06 05:04:03")
	if err != nil {
		t.Errorf("convertTime() error")
	}
	SetTimeZone(tz)
	out2, err := ct("2025-07-06 08:04:03")
	if err != nil {
		t.Errorf("convertTime() error")
	}
	if !out1.Equal(*out2) {
		t.Errorf("convertTime() = %v, want %v", out1, out2)
	}
}
