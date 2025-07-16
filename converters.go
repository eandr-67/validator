package validator

import (
	"fmt"
)

type Converter[T any] func(raw any) (*T, *string)

// Convert базовый преобразователь, используемый для типов данных, автоматически декодируемых из JSON в Go any.
func Convert[T any](raw any) (*T, *string) {
	if raw == nil {
		return nil, nil
	}

	var tmp T
	fmt.Printf("%T %T\n", tmp, raw)

	if v, ok := raw.(T); ok {
		return &v, nil
	} else {
		return &v, &ErrMsg[CodeTypeIncorrect]
	}
}

// convertInt преобразователь, используемых для получения целых чисел,
// т.к. декодер JSON -> any преобразует все числа в float64
// (что полностью соответствует семантике JavaScript).
func convertInt(raw any) (*int64, *string) {
	var t int64
	switch v := raw.(type) {
	case nil:
		return nil, nil
	case int64:
		return &v, nil
	case float64:
		t = int64(v)
		return &t, nil
	}
	return &t, &ErrMsg[CodeTypeIncorrect]
}
