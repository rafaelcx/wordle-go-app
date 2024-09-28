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
		performCheck(gameState)

		if isSolvedGame(gameState) {
			output.PrintGuessHistory(gameState)
			output.PrintSolvedMsg()
			return
		}
		output.PrintGuessHistory(gameState)
	}

	output.PrintGameOverMsg()
}

func performCheck(gameState *state.Game) {
	guess := gameState.GetCurrentGuess()
	answer := gameState.GetAnswer()

	for i := 0; i < len(guess.LetterValue); i++ {
		for j := 0; j < len(answer.LetterValue); j++ {

			lettersMatch := guess.LetterValue[i] == answer.LetterValue[j]

			if lettersMatch && i == j && !answer.IsLetterChecked(j) && !guess.IsLetterChecked(i) {
				guess.SetLetterStateCorrect(i)
				guess.SetLetterAsChecked(i)
				answer.SetLetterAsChecked(j)
				continue
			}

			if lettersMatch && i != j && !answer.IsLetterChecked(j) && !guess.IsLetterChecked(i) {
				guess.SetLetterStatePresent(i)
				guess.SetLetterAsChecked(i)
				answer.SetLetterAsChecked(j)
				continue
			}
		}

		if !guess.IsLetterStateCorrect(i) && !guess.IsLetterStatePresent(i) {
			guess.SetLetterStateWrong(i)
		}
	}

	guess.ResetLetterCheck()
	answer.ResetLetterCheck()
}

func generateAnswer() *state.Word {
	possibleAnswers := file.ParseAsSlice()
	randomIndex := rand.Intn(len(possibleAnswers))
	answerAsStr := strings.ToUpper(possibleAnswers[randomIndex])
	return state.NewWord(answerAsStr)
}

func isSolvedGame(gameState *state.Game) bool {
	for i := 0; i < state.WordSize; i++ {
		if !gameState.GetCurrentGuess().IsLetterStateCorrect(i) {
			return false
		}
	}
	return true
}
