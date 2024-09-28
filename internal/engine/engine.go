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

	gameState.SetAnswer(generateAnswer())

	for gameState.GetAttemptNumber() < maxRounds {
		userInput, err := input.GetInput()
		if err != nil {
			continue
		}

		gameState.SetCurrentGuess(state.NewWord(userInput))

		if gameState.GetCurrentGuess().AsString() == gameState.GetAnswer().AsString() {
			output.PrintGuessHistory(gameState)
			output.PrintSolvedMsg()
			return
		}

		output.PrintGuessHistory(gameState)
	}

	output.PrintGameOverMsg()
}

func generateAnswer() *state.Word {
	possibleAnswers := file.ParseAsSlice()
	randomIndex := rand.Intn(len(possibleAnswers))
	answerAsStr := strings.ToUpper(possibleAnswers[randomIndex])
	return state.NewWord(answerAsStr)
}
