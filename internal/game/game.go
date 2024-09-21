package game

type Game struct {
	attemptNumber int
	currentGuess  string
	guessHistory  []string
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) GetAttemptNumber() int {
	return g.attemptNumber
}

func (g *Game) GetCurrentGuess() string {
	return g.currentGuess
}

func (g *Game) GetGuessHistory() []string {
	if g.currentGuess != "" {
		g.guessHistory = append(g.guessHistory, g.currentGuess)
	}
	return g.guessHistory
}

func (g *Game) SetCurrentGuess(guess string) {
	if g.currentGuess != "" {
		g.guessHistory = append(g.guessHistory, g.currentGuess)
	}
	g.attemptNumber++
	g.currentGuess = guess
}
