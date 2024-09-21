package validator

import "errors"

type wordLengthValidator struct{}

func (validator wordLengthValidator) validate(s string) error {
	if len(s) != 5 {
		return errors.New("your guess should be five letters long, try again")
	}
	return nil
}
