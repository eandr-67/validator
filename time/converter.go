package time

import (
	"time"

	"github.com/eandr-67/validator"
)

// convertTime возвращает замыкание, преобразующее any в указатель на time.Time.
// Преобразование производится в 2 этапа: сначала any преобразуется в string, а потом эта строка декодируется в
// time.Time перебором форматов formats - пока какой-нибудь из форматов не подойдёт.
func convertTime(formats []string) func(any) (*time.Time, *string) {
	return func(raw any) (*time.Time, *string) {
		var res time.Time
		var err error
		s, e := validator.Convert[string](raw)
		if e != nil {
			return &res, e
		} else if s == nil {
			return nil, nil
		}
		for _, format := range formats {
			if res, err = time.ParseInLocation(format, *s, timeZone); err == nil {
				return &res, nil
			}
		}
		return &res, &validator.ErrMsg[validator.CodeFormatIncorrect]
	}
}
