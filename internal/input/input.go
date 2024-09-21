package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() string {
	fmt.Print("Enter your guess: ")
	return getGuess()
}

func getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return guess
}
