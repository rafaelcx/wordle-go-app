package state

import (
	"strings"
)

const cBlack string = "\033[0m"

type Word struct {
	LetterValue []string
	LetterState []string
}

func NewWord(s string) *Word {
	letterValueArray := make([]string, len(s))
	letterStateArray := make([]string, len(s))

	for i, r := range s {
		letterValueArray[i] = string(r)
		letterStateArray[i] = cBlack
	}

	return &Word{
		LetterValue: letterValueArray,
		LetterState: letterStateArray,
	}
}

func (w *Word) AsString() string {
	return strings.Join(w.LetterValue, "")
}
