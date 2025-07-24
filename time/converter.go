package time

import (
	"time"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// timeConverter возвращает замыкание, преобразующее any в указатель на time.Time.
// Преобразование производится в 2 этапа: сначала any преобразуется в string, а потом эта строка декодируется в
// time.Time перебором форматов formats - пока какой-нибудь из форматов не подойдёт.
func timeConverter(formats []string) validator.Converter[time.Time] {
	return func(raw any, err *errs.Errors) *time.Time {
		var e error
		var res time.Time
		switch v := raw.(type) {
		case nil:
			return nil
		case string:
			for _, format := range formats {
				if res, e = time.ParseInLocation(format, v, timeZone); e == nil {
					return &res
				}
			}
			err.Add("", validator.ErrMsg[validator.ErrFormatIncorrect])
		default:
			err.Add("", validator.ErrMsg[validator.ErrTypeIncorrect])

		}
		return &res
	}
}
