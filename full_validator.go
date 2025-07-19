package validator

import (
	"slices"

	"github.com/eandr-67/errs"
)

// NewFull создаёт новый валидатор Full.
// Получает преобразователь, два набора действий, обработчик компонентов и возвращает указатель на созданный валидатор.
func NewFull[T any](converter Converter[T], before, after Rules[T], handler Handler[T]) *Full[T] {
	return &Full[T]{convert: converter, before: slices.Clone(before), after: slices.Clone(after), handler: handler}
}

// Full реализует валидатор, предназначенный для типов сложной структуры,
// требующих валидировать не только значение в целом, но и отдельные компоненты этого значения.
type Full[T any] struct {
	convert       Converter[T]
	before, after Rules[T]
	handler       Handler[T]
}

// Handler определяет интерфейс обработчика компонентов проверяемого значения.
type Handler[T any] interface {
	// Handle реализует обработку компонентов значения.
	// Получает указатель на значение и указатель на список ошибок.
	// Возвращает указатель на значение.
	Handle(*T, *errs.Errors) *T
}

// Do валидатора Full реализует трёхстадийный обработчик указателя на значение, возвращённого преобразователем.
// Сначала к указателю применяется набор действий beforerrs.
// Затем выполняется Handler, реализующий обработку компонентов значения.
// После чего к указателю применяется набор действий after.
func (v *Full[T]) Do(raw any) (res any, err *errs.Errors) {
	err = &errs.Errors{}
	if val, msg := v.convert(raw); msg != nil {
		err.Add("", *msg)
		return nil, err
	} else if val = v.before.Execute(val, "", err); val == nil {
		return nil, err
	} else if val = v.handler.Handle(val, err); val == nil {
		return nil, err
	} else if val = v.after.Execute(val, "", err); val == nil {
		return nil, err
	} else {
		return *val, err
	}
}
