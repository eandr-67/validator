package array

import (
	"slices"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// LenEq проверяет, что длина массива больше или равна заданной.
func LenEq(value int) validator.Action[[]any] {
	return func(val *[]any, key string, err *errs.Errors) (*[]any, bool) {
		if len(*val) == value {
			return val, true
		}
		err.Add(key, validator.ErrMsg[validator.CodeLengthIncorrect])
		return nil, false
	}
}

// LenNe проверяет, что длина массива меньше или равна заданной.
func LenNe(value int) validator.Action[[]any] {
	return func(val *[]any, key string, err *errs.Errors) (*[]any, bool) {
		if len(*val) != value {
			return val, true
		}
		err.Add(key, validator.ErrMsg[validator.CodeLengthIncorrect])
		return nil, false
	}
}

// LenGe проверяет, что длина массива больше или равна заданной.
func LenGe(value int) validator.Action[[]any] {
	return func(val *[]any, key string, err *errs.Errors) (*[]any, bool) {
		if len(*val) > value {
			return val, true
		}
		err.Add(key, validator.ErrMsg[validator.CodeLengthIncorrect])
		return nil, false
	}
}

// LenLe проверяет, что длина массива меньше или равна заданной.
func LenLe(value int) validator.Action[[]any] {
	return func(val *[]any, key string, err *errs.Errors) (*[]any, bool) {
		if len(*val) < value {
			return val, true
		}
		err.Add(key, validator.ErrMsg[validator.CodeLengthIncorrect])
		return nil, false
	}
}

// LenIn проверяет, что длина массива входит в заданный список.
func LenIn(values ...int) validator.Action[[]any] {
	return func(val *[]any, key string, err *errs.Errors) (*[]any, bool) {
		if slices.Contains(values, len(*val)) {
			return val, true
		}
		err.Add(key, validator.ErrMsg[validator.CodeLengthIncorrect])
		return nil, false
	}
}

// LenNotIn проверяет, что длина массива не входит в заданный список.
func LenNotIn(values ...int) validator.Action[[]any] {
	return func(val *[]any, key string, err *errs.Errors) (*[]any, bool) {
		if !slices.Contains(values, len(*val)) {
			return val, true
		}
		err.Add(key, validator.ErrMsg[validator.CodeLengthIncorrect])
		return nil, false
	}
}
