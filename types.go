package validator

// Int получает набор правил валидации и возвращает новый построитель валидатора целого числа.
func Int(rules ...Action[int64]) *Build[int64] {
	return &Build[int64]{
		Convert: convertInt,
		Rules:   rules,
	}
}

// Float получает набор правил валидации и возвращает новый построитель валидатора вещественного числа.
func Float(rules ...Action[float64]) *Build[float64] {
	return &Build[float64]{
		Convert: Convert[float64],
		Rules:   rules,
	}
}

// String получает набор правил валидации и возвращает новый построитель валидатора строки.
func String(rules ...Action[string]) *Build[string] {
	return &Build[string]{
		Convert: Convert[string],
		Rules:   rules,
	}
}

// Bool получает набор правил валидации и возвращает новый построитель валидатора логического значения.
func Bool(rules ...Action[bool]) *Build[bool] {
	return &Build[bool]{
		Convert: Convert[bool],
		Rules:   rules,
	}
}
