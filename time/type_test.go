package time

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:15", timeZone)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:17", timeZone)
	b := Time([]string{"2006-01-02 15:04:05"}, Gt(t1)).Append(Lt(t2))
	if len(b.Rules) != 2 {
		t.Errorf("Rule length should be 2")
	}
	if len(b.formats) != 1 {
		t.Errorf("formats length should be 1")
	}
	_ = b.Validator()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else if v, ok := r.(error); !ok {
			t.Errorf("Panic [%T][%#v] is not a error", r, r)
		} else if v.Error() != "formats cannot be empty" {
			t.Errorf("Panic [%s] is wrong", v.Error())
		}
	}()
	_ = Time([]string{}, Gt(t1)).Append(Lt(t2)).Validator()
}
