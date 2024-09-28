package state

import (
	"strings"
)

const cRed string = "\033[31m"
const cGreen string = "\033[32m"
const cYellow string = "\033[33m"
const cBlack string = "\033[0m"

type Word struct {
	LetterValue []string
	LetterState []string
	LetterCheck []bool
}

func NewWord(s string) *Word {
	letterValueArray := make([]string, len(s))
	letterStateArray := make([]string, len(s))
	letterCheckArray := make([]bool, len(s))

	for i, r := range s {
		letterValueArray[i] = string(r)
		letterStateArray[i] = cBlack
		letterCheckArray[i] = false
	}

	return &Word{
		LetterValue: letterValueArray,
		LetterState: letterStateArray,
		LetterCheck: letterCheckArray,
	}
}

func (w *Word) AsString() string {
	return strings.Join(w.LetterValue, "")
}

func (w *Word) ResetLetterCheck() {
	w.LetterCheck = make([]bool, len(w.LetterCheck))
}

func (w *Word) IsLetterStateCorrect(index int) bool {
	return w.LetterState[index] == cGreen
}

func (w *Word) IsLetterStatePresent(index int) bool {
	return w.LetterState[index] == cYellow
}

func (w *Word) SetLetterStateCorrect(index int) {
	w.LetterState[index] = cGreen
}

func (w *Word) SetLetterStatePresent(index int) {
	w.LetterState[index] = cYellow
}

func (w *Word) SetLetterStateWrong(index int) {
	w.LetterState[index] = cRed
}

func (w *Word) SetLetterAsChecked(index int) {
	w.LetterCheck[index] = true
}

func (w *Word) IsLetterChecked(index int) bool {
	return w.LetterCheck[index]
}
