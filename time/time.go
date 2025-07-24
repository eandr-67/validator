package time

import (
	"time"

	"github.com/eandr-67/validator"
)

// Time фабрика, создающая построитель валидатора time.Time.
// Получает массив допустимых форматов даты/времени и перечисление действий.
func Time(formats []string, actions ...validator.Action[time.Time]) *builder {
	return &builder{actions: actions, formats: formats}
}
