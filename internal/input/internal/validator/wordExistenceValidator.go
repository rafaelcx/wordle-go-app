package validator

import (
	"errors"
	"strings"

	"github.com/rafaelcx/wordle-go-cli/internal/file"
)

type wordExistenceValidator struct{}

func (validator wordExistenceValidator) validate(s string) error {
	var wordList []string = file.ParseAsSlice()

	for i := 0; i < len(wordList); i++ {
		if s == strings.ToUpper(wordList[i]) {
			return nil
		}
	}
	return errors.New("your guess is not a valid word, try again")
}
