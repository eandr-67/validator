package validator

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
	Convert Converter[T]
	Rules[T]
}

// Validator создаёт валидатор Simple из построителя.
func (b *Build[T]) Validator() Validator {
	return NewSimple[T](b.Convert, b.Rules)
}

func (b *Build[T]) Append(rules ...Action[T]) *Build[T] {
	b.Rules.Append(rules...)
	return b
}
