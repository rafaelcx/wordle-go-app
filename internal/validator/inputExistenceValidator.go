package validator

import (
	"errors"
	"strings"

	"github.com/rafaelcx/wordle-app/internal/file"
)

type inputExistenceValidator struct {
	input string
}

func (v inputExistenceValidator) validate() error {
	var entireWordList []string = file.ParseAsSlice()

	for i := 0; i < len(entireWordList); i++ {
		if v.input == strings.ToUpper(entireWordList[i]) {
			return nil
		}
	}
	return errors.New("your guess is not a valid word, try again")
}
