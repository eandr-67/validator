package string

import (
	"regexp"
	"slices"
	"unicode/utf8"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// Regex генератор действия, регистрирующего ошибку ErrFormatIncorrect,
// если значение не соответствует шаблону регулярки pattern.
func Regex(pattern string) validator.Action[string] {
	temp := regexp.MustCompile(pattern)
	return func(value *string, err *errs.Errors) (*string, bool) {
		if temp.MatchString(*value) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrFormatIncorrect])
		return value, false
	}
}

// NotRegex генератор действия, регистрирующего ошибку ErrFormatIncorrect,
// если значение соответствует шаблону регулярки pattern.
func NotRegex(pattern string) validator.Action[string] {
	temp := regexp.MustCompile(pattern)
	return func(value *string, err *errs.Errors) (*string, bool) {
		if !temp.MatchString(*value) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrFormatIncorrect])
		return value, false
	}
}

// LenEq генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения не равна lng.
func LenEq(lng int) validator.Action[string] {
	return func(value *string, err *errs.Errors) (*string, bool) {
		if utf8.RuneCountInString(*value) == lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenNe генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения равна lng.
func LenNe(lng int) validator.Action[string] {
	return func(value *string, err *errs.Errors) (*string, bool) {
		if utf8.RuneCountInString(*value) != lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenGe генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения меньше lng.
func LenGe(lng int) validator.Action[string] {
	return func(value *string, err *errs.Errors) (*string, bool) {
		if utf8.RuneCountInString(*value) >= lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenLe генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения больше lng.
func LenLe(lng int) validator.Action[string] {
	return func(value *string, err *errs.Errors) (*string, bool) {
		if utf8.RuneCountInString(*value) <= lng {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenIn генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения не входит в набор lngs.
func LenIn(lngs ...int) validator.Action[string] {
	return func(value *string, err *errs.Errors) (*string, bool) {
		if slices.Contains(lngs, utf8.RuneCountInString(*value)) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}

// LenNotIn генератор действия, регистрирующего ошибку ErrLengthIncorrect, если длина значения входит в набор lngs.
func LenNotIn(lngs ...int) validator.Action[string] {
	return func(value *string, err *errs.Errors) (*string, bool) {
		if !slices.Contains(lngs, utf8.RuneCountInString(*value)) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrLengthIncorrect])
		return value, false
	}
}
