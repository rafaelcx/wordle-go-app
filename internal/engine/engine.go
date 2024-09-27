package engine

import (
	"math/rand"
	"strings"

	"github.com/rafaelcx/wordle-go-cli/internal/file"
	"github.com/rafaelcx/wordle-go-cli/internal/input"
	"github.com/rafaelcx/wordle-go-cli/internal/output"
	"github.com/rafaelcx/wordle-go-cli/internal/state"
)

const maxRounds int = 6

func Play(gameState *state.Game) {
	output.PrintIntro()

	answer := generateAnswer()

	for gameState.GetAttemptNumber() < maxRounds {
		userInput, err := input.GetInput()
		if err != nil {
			continue
		}

		guess := state.NewWord(userInput)
		gameState.SetCurrentGuess(guess)

		if gameState.GetCurrentGuess().AsString() == answer {
			output.PrintSolvedMsg()
			break
		}
	}

	output.PrintGameOverMsg()
	output.PrintGuessHistory(gameState)
}

func generateAnswer() string {
	possibleAnswers := file.ParseAsSlice()
	randomIndex := rand.Intn(len(possibleAnswers))
	return strings.ToUpper(possibleAnswers[randomIndex])
}
