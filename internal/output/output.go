package output

import (
	"fmt"

	"github.com/rafaelcx/wordle-go-cli/internal/state"
)

func PrintIntro() {
	fmt.Println(" =========================================================== ")
	fmt.Println(" WORDLE GAME MOCK - Play a horrible copy of Wordle for free! ")
	fmt.Println(" ===========================================================\n ")
}

func PrintGuessHistory(gameState *state.Game) {
	guessHitory := gameState.GetGuessHistory()

	for _, value := range guessHitory {
		value.PrintAsString()
		fmt.Println()
	}
}
