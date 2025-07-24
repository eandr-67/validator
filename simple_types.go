package validator

import (
	"github.com/eandr-67/errs"
)

// NewSimpleBuilder создаёт построитель простого валидатора, обрабатывающего значения без учёта их внутренней структуры.
// Получает преобразователь и набор действий.
// Возвращает указатель на созданный построитель.
func NewSimpleBuilder[T any](c Converter[T], a ...Action[T]) *SimpleBuilder[T] {
	return &SimpleBuilder[T]{converter: c, actions: a}
}

// String фабрика, создающая построитель валидатора строки.
func String(actions ...Action[string]) *SimpleBuilder[string] {
	return NewSimpleBuilder(simpleConverter[string], actions...)
}

// Bool фабрика, создающая построитель валидатора логического значения.
func Bool(actions ...Action[bool]) *SimpleBuilder[bool] {
	return NewSimpleBuilder(simpleConverter[bool], actions...)
}

// Float фабрика, создающая построитель валидатора вещественного числа.
func Float(actions ...Action[float64]) *SimpleBuilder[float64] {
	return NewSimpleBuilder(simpleConverter[float64], actions...)
}

// Int фабрика, создающая построитель валидатора целого числа.
func Int(actions ...Action[int64]) *SimpleBuilder[int64] {
	return NewSimpleBuilder(intConverter, actions...)
}

// Any фабрика, создающая построитель валидатора произвольного типа.
func Any(actions ...Action[any]) *SimpleBuilder[any] {
	return NewSimpleBuilder(
		func(raw any, err *errs.Errors) (value *any) {
			if raw == nil {
				return nil
			}
			return &raw
		}, actions...)
}

// SimpleBuilder реализует построитель простого валидатора значения типа T, не содержащего скрытых действий.
type SimpleBuilder[T any] struct {
	converter Converter[T]
	actions   []Action[T]
}

// Compile производит сборку валидатора из установленных параметров построителя
func (sb *SimpleBuilder[T]) Compile() Validator {
	return NewValidator(sb.converter, sb.actions...)
}

// Add добавляет действия в набор.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
func (sb *SimpleBuilder[T]) Add(actions ...Action[T]) *SimpleBuilder[T] {
	sb.actions = append(sb.actions, actions...)
	return sb
}
