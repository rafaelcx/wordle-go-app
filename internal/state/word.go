package state

import "fmt"

const cBlack string = "\033[0m"

type Word struct {
	letterValue []string
	letterState []string
}

func NewWord(s string) *Word {
	letterValueArray := make([]string, len(s))
	letterStateArray := make([]string, len(s))

	for i, r := range s {
		letterValueArray[i] = string(r)
		letterStateArray[i] = cBlack
	}

	return &Word{
		letterValue: letterValueArray,
		letterState: letterStateArray,
	}
}

func (w *Word) PrintAsString() {
	for index, value := range w.letterValue {
		fmt.Print(w.letterState[index] + value + " ")
	}
}
