package game

type Game struct {
	attemptNumber int
	currentGuess  string
	guessHistory  []string
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) SetCurrentGuess(guess string) {
	g.currentGuess = guess
	g.guessHistory = append(g.guessHistory, guess)
	g.attemptNumber++
}

func (g *Game) GetAttemptNumber() int {
	return g.attemptNumber
}

func (g *Game) GetCurrentGuess() string {
	return g.currentGuess
}

func (g *Game) GetGuessHistory() []string {
	return g.guessHistory
}
