package main

import (
	"fmt"

	"github.com/rafaelcx/wordle-go-cli/internal/game"
	"github.com/rafaelcx/wordle-go-cli/internal/input"
)

func main() {
	fmt.Println(" =========================================================== ")
	fmt.Println(" WORDLE GAME MOCK - Play a horrible copy of Wordle for free! ")
	fmt.Println(" ===========================================================\n ")

	var gameState *game.Game = game.NewGame()

	for gameState.GetAttemptNumber() != 6 {
		userInput, err := input.GetInput()
		if err != nil {
			continue
		}

		gameState.SetCurrentGuess(userInput)
	}

	fmt.Println(gameState.GetGuessHistory())
}
