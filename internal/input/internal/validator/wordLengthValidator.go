package validator

import "errors"

type wordLengthValidator struct{}

func (validator wordLengthValidator) validate(s string) error {
	if len(s) != 5 {
		msg := "your guess should be five letters long, try again"
		return errors.New(msg)
	}
	return nil
}
