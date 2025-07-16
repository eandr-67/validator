package validator

import (
	"testing"
)

func Test_init(t *testing.T) {
	var tbl = map[int]string{
		CodeTypeIncorrect:   "type",
		CodeFormatIncorrect: "format",
		CodeLengthIncorrect: "length",
		CodeValueIncorrect:  "value",
		CodeValueIsNull:     "is_null",
		CodeKeyMissed:       "missed",
		CodeKeyUnknown:      "unknown",
	}

	if len(tbl) != len(ErrMsg) {
		t.Errorf("len(ErrMsg) = %d; want %d", len(ErrMsg), len(tbl))
	}

	for k, v := range tbl {
		if ErrMsg[k] != v {
			t.Errorf("ErrMsg[%d] = %s; want %s", k, ErrMsg[k], v)
		}
	}
}
