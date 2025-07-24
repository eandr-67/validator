package uuid

import (
	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
	"github.com/google/uuid"
)

func uuidConverter(raw any, err *errs.Errors) *uuid.UUID {
	var e error
	var res uuid.UUID
	switch v := raw.(type) {
	case nil:
		return nil
	case string:
		if res, e = uuid.Parse(v); e != nil {
			err.Add("", validator.ErrMsg[validator.ErrFormatIncorrect])
		}
	default:
		err.Add("", validator.ErrMsg[validator.ErrTypeIncorrect])

	}
	return &res
}
