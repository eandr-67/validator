package time

import (
	"testing"
	"time"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
	"github.com/stretchr/testify/assert"
)

func TestAction_Compare(t *testing.T) {
	var val [3]time.Time
	val[0], _ = time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:15", timeZone)
	val[1], _ = time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:16", timeZone)
	val[2], _ = time.ParseInLocation("2006-01-02 15:04:05", "2025-07-01 20:15:17", timeZone)
	var err = errs.Errors{}
	var cmp = map[string]struct {
		op  validator.Action[time.Time]
		res [3]bool
	}{
		"Eq":    {Eq(val[1]), [3]bool{false, true, false}},
		"Ne":    {Ne(val[1]), [3]bool{true, false, true}},
		"Lt":    {Lt(val[1]), [3]bool{true, false, false}},
		"Le":    {Le(val[1]), [3]bool{true, true, false}},
		"Gt":    {Gt(val[1]), [3]bool{false, false, true}},
		"Ge":    {Ge(val[1]), [3]bool{false, true, true}},
		"In":    {In(val[1], val[2]), [3]bool{false, true, true}},
		"NotIn": {NotIn(val[1], val[2]), [3]bool{true, false, false}},
	}
	for k, v := range cmp {
		for i, p := range v.res {
			err = nil
			var inp = val[i]
			out, ok := v.op(&inp, &err)
			assert.Equalf(t, *out, inp, k)
			if p {
				assert.Nilf(t, err, k)
				assert.Truef(t, ok, k)
			} else {
				assert.Equalf(t, err, errs.Errors{"": {validator.ErrMsg[validator.ErrValueIncorrect]}}, k)
				assert.Falsef(t, ok, k)
			}
		}
	}
}
