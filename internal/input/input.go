package input

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rafaelcx/wordle-go-cli/internal/input/internal/sanitizer"
)

func GetInput() string {
	fmt.Print("Enter your guess: ")

	guess := getGuess()
	return sanitizer.Sanitize(guess)
}

func getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return guess
}
