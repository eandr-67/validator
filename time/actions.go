package time

import (
	"slices"
	"time"

	"github.com/eandr-67/errs"
	"github.com/eandr-67/validator"
)

func Eq(value time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if elem.Equal(value) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func Ne(value time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !elem.Equal(value) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func In(values ...time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if slices.ContainsFunc(values, elem.Equal) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func NotIn(values ...time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !slices.ContainsFunc(values, elem.Equal) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func Gt(value time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if elem.After(value) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func Ge(value time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !elem.Before(value) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func Lt(value time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if elem.Before(value) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}

func Le(value time.Time) validator.Action[time.Time] {
	return func(elem *time.Time, err *errs.Errors) (*time.Time, bool) {
		if !elem.After(value) {
			return elem, true
		}
		err.Add("", validator.ErrMsg[validator.ErrValueIncorrect])
		return elem, false
	}
}
