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

	// message.PrintSolutionMsg(answer)

	for i := 1; i <= attemptTotal; i++ {
		guess := getGuessInput()

		if guess == answer {
			message.PrintSolvedMsg(i)
			return
		}

		roundState := engine.CalculateState(guess, answer)
		roundStateAsStr := ""
		for k := 0; k < 5; k++ {
			roundStateAsStr = roundStateAsStr + roundState[k]
		}

		history = append(history, roundStateAsStr)
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
