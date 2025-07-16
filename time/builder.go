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

func (build *Build) Validator() validator.Validator {
	if len(build.formats) == 0 {
		panic(errors.New("formats cannot be empty"))
	}
	return validator.NewSimple[time.Time](convertTime(build.formats), build.Rules)
}
