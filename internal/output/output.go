package output

import (
	"fmt"

	"github.com/rafaelcx/wordle-go-cli/internal/state"
)

func PrintIntro() {
	fmt.Println("============================================================ ")
	fmt.Println("WORDLE GAME MOCK - Play a horrible copy of Wordle for free! ")
	fmt.Println("===========================================================\n ")
}

func PrintSolvedMsg() {
	fmt.Println()
	fmt.Println("Solved!!!")
	fmt.Println()
}

func PrintGameOverMsg(gameState *state.Game) {
	fmt.Println()
	fmt.Println("Game Over!!!")
	fmt.Println("The asnwer was: " + gameState.GetAnswer().AsString())
	fmt.Println()
}

func PrintGuessHistory(gameState *state.Game) {
	fmt.Println()
	guessHitory := gameState.GetGuessHistory()

	for _, word := range guessHitory {
		printWordAsString(word)
		fmt.Println()
	}
	fmt.Println(state.ColorBlack)
}

func PrintKeyboard(gameState *state.Game) {
	fmt.Println("Keyboard state:")
	fmt.Println()

	keyboard := gameState.GetKeyboard()
	for i, letter := range keyboard.Order {
		fmt.Print(keyboard.Letters[letter] + letter + " ")
		if i == 10 || i == 19 {
			fmt.Println()
		}
	}

	fmt.Println(state.ColorBlack)
	fmt.Println()
}

func printWordAsString(w *state.Word) {
	for index, value := range w.LetterValue {
		fmt.Print(w.LetterState[index] + value + " ")
	}
}
