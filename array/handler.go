package array

import (
	"fmt"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// handle реализует обработчик элементов массива
type handle struct {
	validator.Validator
}

// Handle вызывается внутри validator.Full для проверки элементов массива
func (h *handle) Handle(data *[]any, err *errs.Errors) *[]any {
	var msg *errs.Errors
	for i, cell := range *data {
		if (*data)[i], msg = h.Do(cell); len(*msg) != 0 {
			err.AddErrors(fmt.Sprintf("[%d]", i), *msg)
		}
	}
	return data
}
