package validator

import (
	"maps"
	"slices"

	"github.com/eandr-67/errs"
)

// Obj фабрика, создающая построитель валидатора значения типа map[string]any (объект JSON)
// Получает на вход действия, добавляемые в начальный набор.
// Возвращает указатель на созданный построитель валидатора объекта.
func Obj(start ...Action[map[string]any]) *objectBuilder {
	return &objectBuilder{
		start:    start,
		fields:   make(map[string]Builder),
		defaults: make(map[string]any),
	}
}

// objectBuilder реализация построителя валидатора объекта.
type objectBuilder struct {
	start, finish []Action[map[string]any]
	fields        map[string]Builder
	required      []string
	defaults      map[string]any
}

// Compile собирает валидатор объекта из пяти компонентов.
func (ob *objectBuilder) Compile() Validator {
	ob.checkNames()
	actions := slices.Clone(ob.start)
	actions = append(
		actions, actionObjectRequired(ob.required), actionObjectDefaults(ob.defaults), actionObjectFields(ob.fields))
	return NewValidator(mapConverter, append(actions, ob.finish...)...)
}

// Start добавляет действия в начальный набор.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
func (ob *objectBuilder) Start(actions ...Action[map[string]any]) *objectBuilder {
	ob.start = append(ob.start, actions...)
	return ob
}

// Finish добавляет действия в конечный набор.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
func (ob *objectBuilder) Finish(actions ...Action[map[string]any]) *objectBuilder {
	ob.finish = append(ob.finish, actions...)
	return ob
}

// Field добавляет в построитель валидатора объекта новое поле с именем name и построителем валидатора builder.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
// Если добавляемое поле уже существует, генерирует панику "field already exists".
func (ob *objectBuilder) Field(name string, builder Builder) *objectBuilder {
	if _, ok := ob.fields[name]; ok {
		panic("field already exists")
	}
	ob.fields[name] = builder
	return ob
}

// FieldList добавляет в построитель валидатора объекта набор полей с именами - ключами fields
// и построителями валидаторов - значениями fields.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
// Если хотя бы одно добавляемое поле уже существует, генерирует панику "field already exists".
func (ob *objectBuilder) FieldList(fields map[string]Builder) *objectBuilder {
	for k, v := range fields {
		ob.Field(k, v)
	}
	return ob
}

// Required объявляет указанные поля обязательными.
// Получает перечисление имён полей.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
func (ob *objectBuilder) Required(fields ...string) *objectBuilder {
	ob.required = append(ob.required, fields...)
	return ob
}

// Default объявляет поле автосоздаваемым. Получает имя поля в name и значение по умолчанию в value.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
// Если поле уже является автосоздаваемым, генерирует панику "default already exists".
func (ob *objectBuilder) Default(name string, value any) *objectBuilder {
	if _, ok := ob.defaults[name]; ok {
		panic("default already exists")
	}
	ob.defaults[name] = value
	return ob
}

// DefaultList объявляет группу полй автосоздаваемыми.
// Ключ fields - имя поля, значение fields - значение по умолчанию.
// Возвращает указатель на построитель для удобной организации цепочек вызовов методов.
// Если хотя бы одно поле уже является автосоздаваемым, генерирует панику "default already exists".
func (ob *objectBuilder) DefaultList(fields map[string]any) *objectBuilder {
	for k, v := range fields {
		ob.Default(k, v)
	}
	return ob
}

// checkNames убирает дубликаты из списка обязательных полей и проверяет,
// что списки обязательных и автосоздаваемых полей не пересекаются и являются подмножествами списка полей.
// Генерирует 3 варианта паники - в зависимости от нарушения.
func (ob *objectBuilder) checkNames() {
	slices.Sort(ob.required)
	ob.required = slices.Compact(ob.required)
	for _, name := range ob.required {
		if _, ok := ob.fields[name]; !ok {
			panic("required field does unknown")
		}
		if _, ok := ob.defaults[name]; ok {
			panic("both required and default")
		}
	}
	for k := range ob.defaults {
		if _, ok := ob.fields[k]; !ok {
			panic("default field does unknown")
		}
	}
}

// actionObjectRequired генератор действия, проверяющего наличие обязательных полей.
func actionObjectRequired(required []string) Action[map[string]any] {
	if len(required) == 0 {
		return nil
	}
	tmp := slices.Clone(required)
	return func(value *map[string]any, err *errs.Errors) (*map[string]any, bool) {
		for _, name := range tmp {
			if _, ok := (*value)[name]; !ok {
				err.Add(name, ErrMsg[ErrKeyMissed])
			}
		}
		return value, true
	}
}

// actionObjectDefaults генератор действия, добавляющего в значение (объект) недостающие автосоздаваемые поля.
func actionObjectDefaults(defaults map[string]any) Action[map[string]any] {
	if len(defaults) == 0 {
		return nil
	}
	tmp := maps.Clone(defaults)
	return func(value *map[string]any, err *errs.Errors) (*map[string]any, bool) {
		for key, def := range tmp {
			if _, ok := (*value)[key]; !ok {
				(*value)[key] = def
			}
		}
		return value, true
	}
}

// actionObjectFields генератор действия, выполняющего вылидаторы для всех полей, существующих в значении (объекте).
func actionObjectFields(fields map[string]Builder) Action[map[string]any] {
	if len(fields) == 0 {
		return nil
	}
	tmp := make(map[string]Validator, len(fields))
	for k, v := range fields {
		if v == nil {
			tmp[k] = Any().Compile()
		} else {
			tmp[k] = v.Compile()
		}
	}
	return func(value *map[string]any, err *errs.Errors) (*map[string]any, bool) {
		var e errs.Errors
		for key, val := range *value {
			if vl, ok := tmp[key]; !ok {
				err.Add(key, ErrMsg[ErrKeyUnknown])
			} else if (*value)[key], e = vl.Do(val); e != nil {
				err.AddErrors(key, e)
			}
		}
		return value, true
	}
}
