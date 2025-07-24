package time

import (
	"time"

	"github.com/eandr-67/validator"
)

// builder реализует построитель валидатора значения time.Time
type builder struct {
	actions []validator.Action[time.Time]
	formats []string
}

// Compile создаёт валидатор time.Time
func (b *builder) Compile() validator.Validator {
	if len(b.formats) == 0 {
		panic("formats cannot be empty")
	}
	return validator.NewValidator(timeConverter(b.formats), b.actions...)
}

// Add добавляет действия в набор действий построителя
func (b *builder) Add(actions ...validator.Action[time.Time]) *builder {
	b.actions = append(b.actions, actions...)
	return b
}
