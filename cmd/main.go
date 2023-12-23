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
)

func main() {
	message.PrintWelcomeMsg()

	const attemptTotal int = 6

	var answer string = generateSolution()
	var history []string

	for i := 1; i <= attemptTotal; i++ {
		guess := getGuessInput()

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
