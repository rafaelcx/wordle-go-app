package engine

func CalculateState(guess, answer string) [5]string {
	const cRed string = "\033[31m"
	const cGreen string = "\033[32m"
	const cYellow string = "\033[33m"
	const cReset string = "\033[0m"

	const checkedFlag string = "checked"
	const wordLen int = 5

	var printStatement [5]string

	guessAsArr := formatIntoArray(guess)
	answerAsArr := formatIntoArray(answer)

	// First we'll check for letters that are green
	for i := 0; i < wordLen; i++ {
		if guessAsArr[i] == answerAsArr[i] && answerAsArr[i] != checkedFlag {
			printStatement[i] = " " + cGreen + guessAsArr[i] + cReset + " "
			answerAsArr[i] = checkedFlag
		}
	}

	// Then we'll check for letter that are yellow
	for i := 0; i < wordLen; i++ {
		for k := 0; k < wordLen; k++ {
			if answerAsArr[k] != checkedFlag && guessAsArr[i] == answerAsArr[k] && printStatement[i] == "" {
				printStatement[i] = " " + cYellow + guessAsArr[i] + cReset + " "
				answerAsArr[k] = checkedFlag
			}
		}
	}

	for i := 0; i < wordLen; i++ {
		if printStatement[i] == "" {
			printStatement[i] = " " + cRed + guessAsArr[i] + cReset + " "
		}
	}

	return printStatement
}

func formatIntoArray(word string) [5]string {
	var wordAsArr [5]string

	for i := 0; i < 5; i++ {
		wordAsArr[i] = string(word[i])
	}
	return wordAsArr
}
