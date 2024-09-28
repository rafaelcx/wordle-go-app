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

func PrintSolvedMsg() {
	fmt.Println()
	fmt.Println("Solved!!!")
}

func PrintGameOverMsg() {
	fmt.Println()
	fmt.Println("Game Over!!!")
}

func PrintGuessHistory(gameState *state.Game) {
	guessHitory := gameState.GetGuessHistory()

	for _, word := range guessHitory {
		printWordAsString(word)
		fmt.Println()
	}
	fmt.Println("\033[0m")
}

func printWordAsString(w *state.Word) {
	for index, value := range w.LetterValue {
		fmt.Print(w.LetterState[index] + value + " ")
	}
}
