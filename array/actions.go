package array

import (
	"slices"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// LenEq генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения не равна lng.
func LenEq(lng int) validator.Action[[]any] {
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		if len(*value) == lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenNe генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения равна lng.
func LenNe(lng int) validator.Action[[]any] {
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		if len(*value) != lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenGe генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения меньше lng.
func LenGe(lng int) validator.Action[[]any] {
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		if len(*value) >= lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenLe генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения больше lng.
func LenLe(lng int) validator.Action[[]any] {
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		if len(*value) <= lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenIn генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения не входит в набор lngs.
func LenIn(lngs ...int) validator.Action[[]any] {
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		if slices.Contains(lngs, len(*value)) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenNotIn генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения входит в набор lngs.
func LenNotIn(lngs ...int) validator.Action[[]any] {
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		if !slices.Contains(lngs, len(*value)) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}
