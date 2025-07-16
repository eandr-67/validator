package object

import (
	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

type handle map[string]validator.Validator

func (h *handle) Handle(data *map[string]any, err *errs.Errors) *map[string]any {
	var msg *errs.Errors
	for key, value := range *data {
		if vl, ok := (*h)[key]; !ok {
			err.Add("."+key, validator.ErrMsg[validator.CodeKeyUnknown])
		} else if (*data)[key], msg = vl.Do(&value); len(*msg) != 0 {
			err.AddErrors("."+key, *msg)
		}
	}
	return data
}
