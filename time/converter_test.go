package time

import (
	"testing"
	"time"

	"github.com/eandr-67/validator"
)

func TestConvertTime(t *testing.T) {
	cnv := convertTime(Default)

	out, err := cnv(nil)
	if err != nil {
		t.Error("err should be nil")
	}
	if out != nil {
		t.Error("out should be nil")
	}

	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:16", timeZone)
	out, err = cnv("2025-07-01 23:15:16+03:00")
	if err != nil {
		t.Error("err should be nil")
	}
	if out == nil {
		t.Error("out should not be nil")
	} else {
		if !out.Equal(tmp) {
			t.Error("tmp should be equal")
		}
	}

	tmp = time.Time{}
	out, err = cnv("2025-07-01 23:15")
	if err == nil {
		t.Error("err should not be nil")
	} else if *err != validator.ErrMsg[validator.CodeFormatIncorrect] {
		t.Error("err should be validator.ErrMsg[validator.CodeFormatIncorrect]")
	}
	if !out.Equal(tmp) {
		t.Error("tmp should be equal")
	}

	out, err = cnv(25)
	if err == nil {
		t.Error("err should not be nil")
	} else if *err != validator.ErrMsg[validator.CodeTypeIncorrect] {
		t.Error("err should be validator.ErrMsg[validator.CodeTypeIncorrect]")
	}
	if !out.Equal(tmp) {
		t.Error("tmp should be equal")
	}
}
