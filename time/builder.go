package time

import (
	"errors"
	"time"

	"github.com/eandr-67/validator"
)

type Build struct {
	validator.Rules[time.Time]
	formats []string
}

func (b *Build) Validator() validator.Validator {
	if len(b.formats) == 0 {
		panic(errors.New("formats cannot be empty"))
	}
	return validator.NewSimple[time.Time](convertTime(b.formats), b.Rules)
}

func (b *Build) Append(rules ...validator.Action[time.Time]) *Build {
	b.Rules.Append(rules...)
	return b
}
