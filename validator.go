package validator

import (
	"fmt"

	"github.com/eandr-67/errs"
)

// NewValidator создаёт валидатор.
// Получает на вход преобразователь и набор действий и возвращает валидатор.
// В валидатор включаются только действия, не равные nil.
func NewValidator[T any](converter Converter[T], actions ...Action[T]) Validator {
	if converter == nil {
		panic("converter cannot be nil")
	}
	v := validator[T]{convert: converter}
	for _, action := range actions {
		if action != nil {
			v.actions = append(v.actions, action)
		}
	}
	return v
}

// Validator определяет нетипизированный (не содержит информацию о типе проверяемого значения) интерфейс,
// реализуемый каждым валидатором.
// Получает на вход сырые данные типа any и возвращает обработанные данные типа any и регистратор ошибок.
// Если ошибок нет, в err возвращается nil.
type Validator interface {
	// Do выполняет собственно процесс валидации.
	Do(raw any) (result any, err errs.Errors)
}

// Builder определяет интерфейс, реализуемый построителем.
type Builder interface {
	// Compile генерирует из построителя нетипизированный валидатор.
	Compile() Validator
}

// validator универсальная, надеюсь, реализация валидатора.
//
// Процесс валидации представляет собой конвейер, запускаемый преобразователем, получающим значение any и
// возвращающим указатель на типизированное значение, и состоящий из набора последовательно выполняемых действий,
// получающих указатель на типизированное значение и возвращающих указатель на преобразованное значение.
// Действия могут менять значения, добавлять поля в объекте и т.д., но не могут изменить тип значения.
type validator[T any] struct {
	convert Converter[T] // convert - преобразователь
	actions []Action[T]  // actions - набор действий
}

// Do реализует процесс валидации, состоящий из выполнения сначала преобразователя, получающего any и возвращающего *T,
// а потом конвейера действий, обрабатывающих возвращённое преобразователем значение.
func (v validator[T]) Do(raw any) (result any, err errs.Errors) {

	defer func() {
		if r := recover(); r != nil {
			err.Add("", fmt.Sprintf(ErrMsg[ErrPanic], r))
		}
	}()

	var value *T
	var cp bool
	if value = v.convert(raw, &err); err != nil {
		return *value, err
	}
	for _, action := range v.actions {
		if value, cp = action(value, &err); !cp {
			break
		}
	}
	if value == nil {
		return nil, err
	}
	return *value, err
}

// Converter определяет преобразователь, получающий значение типа any
// и возвращающий указатель на значение заданного типа и регистратор ошибок.
//
// Ошибка преобразования регистрируется с ключом "".
type Converter[T any] func(raw any, err *errs.Errors) (value *T)

// Action определяет действие: минимальную единицу валидатора, проверяющую одно условие и, возможно,
// меняющую проверяемое значение.
// Получает указатель на проверяемое значение и указатель на регистратор ошибок.
// Возвращает новое значение и флаг продолжения выполнения набора действий.
//
// Все импортируемые действия, включенные в пакет validator, регистрируют ошибки с ключом "".
type Action[T any] func(value *T, err *errs.Errors) (newValue *T, continueProcessing bool)
