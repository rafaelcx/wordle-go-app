package validator

import (
	"errors"
)

type inputSizeValidator struct {
	input string
}

func (v inputSizeValidator) validate() error {
	if len(v.input) != 5 {
		return errors.New("your guess should be five letters long, try again")
	}
	return nil
}
