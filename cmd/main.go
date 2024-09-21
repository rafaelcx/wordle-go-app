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
		gameState.SetCurrentGuess(input.GetInput())
	}

	fmt.Println(gameState.GetGuessHistory())
}
