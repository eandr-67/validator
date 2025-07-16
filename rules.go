package validator

import (
	"github.com/eandr-67/errs"
)

// Rules используется для хранения набора действий, применяемых к обрабатываемому значению.
type Rules[T any] []Action[T]

// Execute применяет набор своих действий к переданному указателю на значение.
// Получает указатель на значение, ключ значения для регистрации в списке ошибок и указатель на список ошибок.
// Возвращает указатель на обработанное значение.
func (r Rules[T]) Execute(elem *T, key string, err *errs.Errors) *T {
	for _, v := range r {
		var ok bool
		if elem, ok = v(elem, key, err); !ok {
			break
		}
	}
	return elem
}

// Append добавляет указанные действия в набор. Действия будут применяться строго в порядке добавления.
func (r *Rules[T]) Append(rules ...Action[T]) *Rules[T] {
	*r = append(*r, rules...)
	return r
}
