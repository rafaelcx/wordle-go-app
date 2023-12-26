package morpher

import "strings"

type inputToUpperCaseMorpher struct {
	input string
}

func (m inputToUpperCaseMorpher) morph() string {
	return strings.ToUpper(m.input)
}
