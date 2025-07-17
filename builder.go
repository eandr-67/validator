package validator

import (
	"errors"
)

func NewBuilder[T any](convert Converter[T], rules Rules[T]) *Build[T] {
	if convert == nil {
		panic(errors.New("convert cannot be nil"))
	}
	if rules == nil {
		rules = Rules[T]{}
	}
	return &Build[T]{convert: convert, rules: rules}
}

// Builder реализуется построителями и определяет сигнатуру метода, генерирующего валидаторы
type Builder interface {
	// Validator метод, которым построитель генерирует готовый к использованию валидатор
	Validator() Validator
}

// Build реализует простой построитель, достаточный для типов данных,
// не имеющих дополнительно валидируемой внутренней структуры (числа, логические значения, строки и т.п.)
// Состоит из преобразователя, возвращающего указатель на валидируемое значение, и набора действий,
// применяемы к этому значению
type Build[T any] struct {
	convert Converter[T]
	rules   Rules[T]
}

// Validator создаёт валидатор Simple из построителя.
func (b *Build[T]) Validator() Validator {
	return NewSimple[T](b.convert, b.rules)
}

func (b *Build[T]) Append(rules ...Action[T]) *Build[T] {
	b.rules.Append(rules...)
	return b
}
