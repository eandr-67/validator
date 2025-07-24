package validator

import (
	"cmp"
	"slices"

	"github.com/eandr-67/errs"
)

// Null действие, останавливающее конвейер, если указатель на значение равен nil. Не регистрирует ошибки.
//
// Если значение - nil, дальнейшие проверки лишены смысла и лишь приведут к ошибкам.
func Null[T any](value *T, _ *errs.Errors) (*T, bool) {
	return value, value != nil
}

// IfNull генератор действия, возвращающего указатель на replace, если указатель на значение равен nil.
// Не регистрирует ошибки.
func IfNull[T any](replace T) Action[T] {
	return func(value *T, _ *errs.Errors) (*T, bool) {
		if value == nil {
			return &replace, true
		}
		return value, true
	}
}

// NotNull действие, регистрирующее ошибку ErrValueIsNull, если указатель на значение равен nil.
func NotNull[T any](value *T, err *errs.Errors) (*T, bool) {
	if value == nil {
		err.Add("", ErrMsg[ErrValueIsNull])
		return nil, false
	}
	return value, true
}

// Eq генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не равно cmp.
func Eq[T comparable](cmp T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if *value == cmp {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// Ne генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение равно cmp.
func Ne[T comparable](cmp T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if *value != cmp {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// In генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не входит в набор cmps.
func In[T comparable](cmps ...T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if slices.Contains(cmps, *value) {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// NotIn генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение входит в набор cmps.
func NotIn[T comparable](cmps ...T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if !slices.Contains(cmps, *value) {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// Gt генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не больше cmp.
func Gt[T cmp.Ordered](cmp T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if *value > cmp {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// Ge генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение меньше cmp.
func Ge[T cmp.Ordered](cmp T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if *value >= cmp {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// Lt генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение не меньше cmp.
func Lt[T cmp.Ordered](cmp T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if *value < cmp {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}

// Le генератор действия, регистрирующего ошибку ErrValueIncorrect, если значение больше cmp.
func Le[T cmp.Ordered](cmp T) Action[T] {
	return func(value *T, err *errs.Errors) (*T, bool) {
		if *value <= cmp {
			return value, true
		}
		err.Add("", ErrMsg[ErrValueIncorrect])
		return value, false
	}
}
