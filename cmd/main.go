package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/rafaelcx/wordle-app/internal/engine"
	"github.com/rafaelcx/wordle-app/internal/file"
	"github.com/rafaelcx/wordle-app/internal/message"
	"github.com/rafaelcx/wordle-app/internal/morpher"
	"github.com/rafaelcx/wordle-app/internal/validator"
)

func main() {
	message.PrintWelcomeMsg()

	const attemptTotal int = 6

	var answer string = strings.ToUpper(generateSolution())
	var history []string

	for i := 1; i <= attemptTotal; i++ {
		guess := getGuessInput()
		guess = morpher.InputToUpperCase(guess)

		err := validator.Validate(guess)
		if err != nil {
			fmt.Println(err.Error())
			i--
			continue
		}

		if guess == answer {
			history = calculateStateAndAppendToHistory(guess, answer, history)
			message.PrintSolvedMsg(i, history)
			return
		}

		history = calculateStateAndAppendToHistory(guess, answer, history)
		message.PrintFailureMsg(history)
	}
	message.PrintGameOverMsg(answer)
}

func generateSolution() string {
	var entireWordList []string = file.ParseAsSlice()
	var entireWordListLen int = len(entireWordList) - 1
	randIndex := rand.Intn(entireWordListLen)
	return entireWordList[randIndex]
}

func getGuessInput() string {
	fmt.Print("Enter your guess: ")
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return strings.TrimSuffix(guess, "\n")
}

func formatRoundStateAsStr(roundState [5]string) string {
	var roundStateAsStr string
	for k := 0; k < len(roundState); k++ {
		roundStateAsStr = roundStateAsStr + roundState[k]
	}
	return roundStateAsStr
}

func calculateStateAndAppendToHistory(guess, answer string, history []string) []string {
	roundState := engine.CalculateState(guess, answer)
	roundStateAsStr := formatRoundStateAsStr(roundState)
	history = append(history, roundStateAsStr)
	return history
}
