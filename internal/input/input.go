package input

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rafaelcx/wordle-go-cli/internal/input/internal/sanitizer"
	"github.com/rafaelcx/wordle-go-cli/internal/input/internal/validator"
)

func GetInput() (string, error) {
	fmt.Print("Enter your guess: ")

	guess := getGuess()
	guess = sanitizer.Sanitize(guess)

	validationErr := validator.Validate(guess)
	if validationErr != nil {
		return "", validationErr
	}
	return guess, nil
}

func getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return guess
}
