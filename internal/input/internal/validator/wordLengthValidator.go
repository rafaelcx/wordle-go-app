package validator

import (
	"errors"

	"github.com/rafaelcx/wordle-go-cli/internal/state"
)

type wordLengthValidator struct{}

func (validator wordLengthValidator) validate(s string) error {
	if len(s) != state.WordSize {
		return errors.New("your guess should be five letters long, try again")
	}
	return nil
}
