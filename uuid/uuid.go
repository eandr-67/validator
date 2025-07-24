package uuid

import (
	"github.com/eandr-67/validator"
	"github.com/google/uuid"
)

func UUID(actions ...validator.Action[uuid.UUID]) *validator.SimpleBuilder[uuid.UUID] {
	return validator.NewSimpleBuilder(uuidConverter, actions...)
}
