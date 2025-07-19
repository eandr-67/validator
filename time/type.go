package time

import (
	"time"

	"github.com/eandr-67/validator"
)

// Time создаёт построитель валидатора time.Time.
// Получает массив допустимых форматов даты/времени и набор действий.
func Time(formats []string, rules ...validator.Action[time.Time]) *Build {
	return &Build{
		Rules:   rules,
		formats: formats,
	}
}
