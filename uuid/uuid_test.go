package uuid

import (
	"testing"

	"github.com/eandr-67/validator"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTime_OK(t *testing.T) {
	u, _ := uuid.Parse("12345678-9abc-def0-1234-56789abcdef0")
	b := UUID(validator.Ne(u))

	assert.NotPanics(t, func() { _ = b.Compile() })

}
