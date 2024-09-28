package state

type Game struct {
	attemptNumber int
	answer        *Word
	currentGuess  *Word
	guessHistory  []*Word
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) SetAnswer(answer *Word) {
	g.answer = answer
}

func (g *Game) SetCurrentGuess(guess *Word) {
	g.currentGuess = guess
	g.guessHistory = append(g.guessHistory, guess)
	g.attemptNumber++
}

func (g *Game) GetAttemptNumber() int {
	return g.attemptNumber
}

func (g *Game) GetAnswer() *Word {
	return g.answer
}

func (g *Game) GetCurrentGuess() *Word {
	return g.currentGuess
}

func (g *Game) GetGuessHistory() []*Word {
	return g.guessHistory
}
