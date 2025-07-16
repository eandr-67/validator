package time

import (
	"time"

	"github.com/eandr-67/validator"
)

func Time(formats []string, rules ...validator.Action[time.Time]) *Build {
	return &Build{
		Rules:   rules,
		formats: formats,
	}
}
