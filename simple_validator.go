package validator

import (
	"slices"

	"github.com/eandr-67/errs"
)

// NewSimple создаёт новый валидатор Simple.
// Получает преобразователь и набор действий и возвращает указатель на созданный валидатор.
func NewSimple[T any](converter Converter[T], rules Rules[T]) *Simple[T] {
	return &Simple[T]{convert: converter, rules: slices.Clone(rules)}
}

// Simple реализует простой валидатор для простых типов, не требующих дополнительной обработки компонентов значения.
type Simple[T any] struct {
	convert Converter[T]
	rules   Rules[T]
}

// Do валидатора Simple реализует простейший вариант обработки: применение заданного набора действий
// к указателю на значение, возвращённому преобразователем.
func (v *Simple[T]) Do(raw any) (res any, err *errs.Errors) {
	err = &errs.Errors{}
	if val, msg := v.convert(raw); msg != nil {
		err.Add("", *msg)
		return nil, err
	} else if val = v.rules.Execute(val, "", err); val == nil {
		return nil, err
	} else {
		return *val, err
	}
}
