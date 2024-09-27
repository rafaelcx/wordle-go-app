package main

import (
	"github.com/rafaelcx/wordle-go-cli/internal/input"
	"github.com/rafaelcx/wordle-go-cli/internal/output"
	"github.com/rafaelcx/wordle-go-cli/internal/state"
)

func main() {
	output.PrintIntro()

	var gameState *state.Game = state.NewGame()

	for gameState.GetAttemptNumber() != 6 {
		userInput, err := input.GetInput()
		if err != nil {
			continue
		}

		guess := state.NewWord(userInput)
		gameState.SetCurrentGuess(guess)
	}

	output.PrintGuessHistory(gameState)
}
