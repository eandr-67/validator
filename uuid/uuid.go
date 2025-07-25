package uuid

import (
	"github.com/eandr-67/validator"
	"github.com/google/uuid"
)

// UUID фабрика, создающая построитель валидатора uuid.UUID.
// Ничем не отличается от фабрик простых построителей пакета validator.
func UUID(actions ...validator.Action[uuid.UUID]) *validator.SimpleBuilder[uuid.UUID] {
	return validator.NewSimpleBuilder(uuidConverter, actions...)
}
