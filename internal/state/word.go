package state

import (
	"strings"
)

const WordSize = 5

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
		letterStateArray[i] = ColorBlack
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
	return w.LetterState[index] == ColorGreen
}

func (w *Word) IsLetterStatePresent(index int) bool {
	return w.LetterState[index] == ColorYellow
}

func (w *Word) SetLetterStateCorrect(index int) {
	w.LetterState[index] = ColorGreen
}

func (w *Word) SetLetterStatePresent(index int) {
	w.LetterState[index] = ColorYellow
}

func (w *Word) SetLetterStateWrong(index int) {
	w.LetterState[index] = ColorRed
}

func (w *Word) SetLetterAsChecked(index int) {
	w.LetterCheck[index] = true
}

func (w *Word) IsLetterChecked(index int) bool {
	return w.LetterCheck[index]
}
