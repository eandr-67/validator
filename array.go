package validator

import (
	"slices"
	"strconv"

	"github.com/eandr-67/errs"
)

// Arr фабрика, создающая построитель валидатора типа []any (массив JSON).
// Получает на вход построитель валидатора, применяемый к элементам массива, и действия, добавляемые в начальный набор.
// Возвращает указатель на созданный построитель валидатора массива.
//
// Если построитель валидатора элементов не задан (cell == nil), проверка/преобразование элементов массива
// не производится.
func Arr(cell Builder, start ...Action[[]any]) *arrayBuilder {
	return &arrayBuilder{cell: cell, start: start}
}

// arrayBuilder реализация построителя валидатора массива.
type arrayBuilder struct {
	start, finish []Action[[]any]
	cell          Builder
}

// Compile создаёт валидатор значения типа []any и возвращает его как нетипизированный интерфейс Validator.
//
// Конвейер валидатора массива состоит из 3 частей:
// начальный набор действий -> скрытое действие обработки элементов массива -> конечны набор действий.
func (ab *arrayBuilder) Compile() Validator {
	b := append(slices.Clone(ab.start), actionArrayCell(ab.cell))
	return NewValidator(simpleConverter[[]any], append(b, ab.finish...)...)
}

// Start добавляет действия в начальный набор.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
func (ab *arrayBuilder) Start(actions ...Action[[]any]) *arrayBuilder {
	ab.start = append(ab.start, actions...)
	return ab
}

// Finish добавляет действия в конечный набор.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
func (ab *arrayBuilder) Finish(actions ...Action[[]any]) *arrayBuilder {
	ab.finish = append(ab.finish, actions...)
	return ab
}

// actionArrayCell генератор действия, применяющего валидатор cell к каждому элементу массива.
func actionArrayCell(cell Builder) Action[[]any] {
	if cell == nil {
		return nil
	}
	vl := cell.Compile()
	return func(value *[]any, err *errs.Errors) (*[]any, bool) {
		var e errs.Errors
		for i, val := range *value {
			if (*value)[i], e = vl.Do(val); e != nil {
				err.AddErrors(strconv.Itoa(i), e)
			}
		}
		return value, true
	}
}
