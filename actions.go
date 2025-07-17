package validator

import (
	"cmp"
	"errors"
	"regexp"
	"slices"

	"github.com/eandr-67/errs"
)

// Action определяет структуру правила валидации (действия).
// Получает указатель на значение, ключ значения для регистрации в списке ошибок и указатель на список ошибок.
// Возвращает указатель на обработанное значение и флаг продолжения валидации.
// В случае ошибки добавляет её текст в список ошибок - обычно, с заданным ключом значения.
type Action[T any] func(elem *T, key string, err *errs.Errors) (*T, bool)

// Null завершает валидацию элемента без установки ошибок, если указатель на элемент равен nil
func Null[T any](elem *T, _ string, _ *errs.Errors) (*T, bool) {
	return elem, elem != nil
}

// IfNull получает значение value и возвращает действие,
// возвращающее value вместо валидируемого значения, если указатель на валидируемое значение равен nil
func IfNull[T any](value T) Action[T] {
	return func(elem *T, _ string, _ *errs.Errors) (*T, bool) {
		if elem == nil {
			return &value, true
		}
		return elem, true
	}
}

// NotNull проверяет, что указатель на валидируемое значение не равен nil.
// Устанавливает ошибку CodeValueIsNull
func NotNull[T any](elem *T, key string, err *errs.Errors) (*T, bool) {
	if elem == nil {
		err.Add(key, ErrMsg[CodeValueIsNull])
		return elem, false
	}
	return elem, true
}

// Eq Применяется только к типам, для которых определены операции == и !=.
// Получает значение value и возвращает действие, проверяющее равенство value и валидируемого значения.
// Если значения не равны, действие добавляет ошибку ValueIncorrect и останавливает валидацию значения.
func Eq[T comparable](value T) Action[T] {
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if *elem == value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// Ne Применяется только к типам, для которых определены операции == и !=.
// Получает значение value и возвращает действие, проверяющее неравенство value и валидируемого значения.
// Если значения равны, действие добавляет ошибку ValueIncorrect и останавливает валидацию значения.
func Ne[T comparable](value T) Action[T] {
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if *elem != value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// In Применяется только к типам, для которых определены операции == и !=.
// Проверяет, что валидируемое значение входит в заданный список.
func In[T comparable](values ...T) Action[T] {
	if len(values) == 0 {
		panic(errors.New("must have at least one value"))
	}
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if slices.Contains(values, *elem) {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// NotIn Применяется только к типам, для которых определены операции == и !=.
// Проверяет, что валидируемое значение не входит в заданный список.
func NotIn[T comparable](values ...T) Action[T] {
	if len(values) == 0 {
		panic(errors.New("must have at least one value"))
	}
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if !slices.Contains(values, *elem) {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// Gt применяется только к типам, значения которых можно упорядочить.
// Проверяет, что валидируемое значение больше заданного
func Gt[T cmp.Ordered](value T) Action[T] {
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if *elem > value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// Ge применяется только к типам, значения которых можно упорядочить.
// Проверяет, что валидируемое значение больше или равно заданному
func Ge[T cmp.Ordered](value T) Action[T] {
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if *elem >= value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// Lt применяется только к типам, значения которых можно упорядочить.
// Проверяет, что валидируемое значение меньше заданного
func Lt[T cmp.Ordered](value T) Action[T] {
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if *elem < value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// Le применяется только к типам, значения которых можно упорядочить.
// Проверяет, что валидируемое значение меньше или равно заданному
func Le[T cmp.Ordered](value T) Action[T] {
	return func(elem *T, key string, err *errs.Errors) (*T, bool) {
		if *elem <= value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeValueIncorrect])
		return elem, false
	}
}

// Regex применяется только к строкам. Проверяет, что строка соответствует заданному регулярному выражению.
func Regex(pattern string) Action[string] {
	temp := regexp.MustCompile(pattern)
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if temp.MatchString(*elem) {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeFormatIncorrect])
		return elem, false
	}
}

// NotRegex применяется только к строкам. Проверяет, что строка не соответствует заданному регулярному выражению.
func NotRegex(pattern string) Action[string] {
	temp := regexp.MustCompile(pattern)
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if !temp.MatchString(*elem) {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeFormatIncorrect])
		return elem, false
	}
}

// LenEq применяется только к строкам. Проверяет, что длина строки равна заданной.
func LenEq(value int) Action[string] {
	if value < 0 {
		panic(errors.New("the value must be not less than 0"))
	}
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if len(*elem) == value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeLengthIncorrect])
		return elem, false
	}
}

// LenNe применяется только к строкам. Проверяет, что длина строки не равна заданной.
func LenNe(value int) Action[string] {
	if value < 0 {
		panic(errors.New("the value must be not less than 0"))
	}
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if len(*elem) != value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeLengthIncorrect])
		return elem, false
	}
}

// LenGe применяется только к строкам. Проверяет, что длина строки больше или равна заданной.
func LenGe(value int) Action[string] {
	if value < 0 {
		panic(errors.New("the value must be not less than 0"))
	}
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if len(*elem) >= value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeLengthIncorrect])
		return elem, false
	}
}

// LenLe применяется только к строкам. Проверяет, что длина строки меньше или равна заданной.
func LenLe(value int) Action[string] {
	if value < 0 {
		panic(errors.New("the value must be not less than 0"))
	}
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if len(*elem) <= value {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeLengthIncorrect])
		return elem, false
	}
}

// LenIn применяется только к строкам. Проверяет, что длина проверяемой строки входит в заданный список.
func LenIn(values ...int) Action[string] {
	if len(values) == 0 {
		panic(errors.New("must have at least one value"))
	}
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if slices.Contains(values, len(*elem)) {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeLengthIncorrect])
		return elem, false
	}
}

// LenNotIn применяется только к строкам. Проверяет, что длина проверяемой строки не входит в заданный список.
func LenNotIn(values ...int) Action[string] {
	if len(values) == 0 {
		panic(errors.New("must have at least one value"))
	}
	return func(elem *string, key string, err *errs.Errors) (*string, bool) {
		if !slices.Contains(values, len(*elem)) {
			return elem, true
		}
		err.Add(key, ErrMsg[CodeLengthIncorrect])
		return elem, false
	}
}
