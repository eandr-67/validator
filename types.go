package validator

// Int получает набор правил валидации и возвращает новый построитель валидатора целого числа.
func Int(rules ...Action[int64]) *Build[int64] {
	return NewBuilder(convertInt, rules)
}

// Float получает набор правил валидации и возвращает новый построитель валидатора вещественного числа.
func Float(rules ...Action[float64]) *Build[float64] {
	return NewBuilder(Convert[float64], rules)
}

// String получает набор правил валидации и возвращает новый построитель валидатора строки.
func String(rules ...Action[string]) *Build[string] {
	return NewBuilder(Convert[string], rules)
}

// Bool получает набор правил валидации и возвращает новый построитель валидатора логического значения.
func Bool(rules ...Action[bool]) *Build[bool] {
	return NewBuilder[bool](Convert[bool], rules)
}
