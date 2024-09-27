package main

import (
	"github.com/rafaelcx/wordle-go-cli/internal/engine"
	"github.com/rafaelcx/wordle-go-cli/internal/state"
)

func main() {
	gameState := state.NewGame()
	engine.Play(gameState)
}
