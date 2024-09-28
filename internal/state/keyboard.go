package state

type Keyboard struct {
	Letters map[string]string
	Order   []string
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		Letters: map[string]string{
			"A": ColorBlack, "B": ColorBlack, "C": ColorBlack, "D": ColorBlack,
			"E": ColorBlack, "F": ColorBlack, "G": ColorBlack, "H": ColorBlack,
			"I": ColorBlack, "J": ColorBlack, "K": ColorBlack, "L": ColorBlack,
			"M": ColorBlack, "N": ColorBlack, "O": ColorBlack, "P": ColorBlack,
			"Q": ColorBlack, "R": ColorBlack, "S": ColorBlack, "T": ColorBlack,
			"U": ColorBlack, "V": ColorBlack, "W": ColorBlack, "X": ColorBlack,
			"Y": ColorBlack, "Z": ColorBlack,
		},
		Order: []string{
			"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P",
			"A", "S", "D", "F", "G", "H", "J", "K", "L",
			"Z", "X", "C", "V", "B", "N", "M",
		},
	}
}

func (k *Keyboard) setLetterColor(letter, color string) {
	k.Letters[letter] = color
}
