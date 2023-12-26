package morpher

func Morph(input string) string {
	for _, morpher := range getMorpherList(input) {
		input = morpher.morph()
	}
	return input
}

func getMorpherList(input string) []Morpher {
	var morphers []Morpher
	morphers = append(morphers, inputToUpperCaseMorpher{input: input})
	return morphers
}
