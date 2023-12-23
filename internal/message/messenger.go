package message

import "fmt"

func PrintWelcomeMsg() {
	fmt.Println(" ========================================================== ")
	fmt.Println(" WORDLE GAME MOCK - Play a Wordle game horrible copy for free! ")
	fmt.Println(" ==========================================================\n ")
}

func PrintGameOverMsg(answer string) {
	fmt.Println("\n\nYou failed miserably!")
	fmt.Printf("The correct answer was %v", answer)
}

func PrintFailureMsg(history []string) {
	fmt.Println("Incorrect - Your history:")
	for i := 0; i < len(history); i++ {
		fmt.Println(history[i])
	}
	fmt.Println()
}

func PrintSolvedMsg(attempts int, history []string) {
	fmt.Println("\nSolved!!!")
	for i := 0; i < len(history); i++ {
		fmt.Println(history[i])
	}
	fmt.Printf("\nAnd it only took you %v times, dummy", attempts)
}

func PrintSolutionMsg(answer string) {
	fmt.Printf("Solution is %v\n\n", answer)
}
