package time

import (
	"slices"
	"time"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

// Eq генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не равно cmp.
func Eq(cmp time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if value.Equal(cmp) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// Ne генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение равно cmp.
func Ne(cmp time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !value.Equal(cmp) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// In генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не входит в набор cmps.
func In(cmps ...time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if slices.ContainsFunc(cmps, value.Equal) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// NotIn генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение входит в набор cmps.
func NotIn(cmps ...time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !slices.ContainsFunc(cmps, value.Equal) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// Gt генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не больше cmp.
func Gt(cmp time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if value.After(cmp) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// Ge генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение меньше cmp.
func Ge(cmp time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !value.Before(cmp) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// Lt генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не меньше cmp.
func Lt(cmp time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if value.Before(cmp) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}

// Le генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение больше cmp.
func Le(cmp time.Time) validator.Action[time.Time] {
	return func(value *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !value.After(cmp) {
			return value, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return value, false
	}
}
