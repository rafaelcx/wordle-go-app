# Wordle Game - CLI version

## Introduction

Wordle is a fun and addictive word puzzle game where players have six attempts to guess a secret five-letter word. After each guess, the game provides feedback on the letters' correctness: 
- A letter in the correct position is highlighted in green.
- A letter in the word but in the wrong position is highlighted in yellow.
- A letter not in the word is highlighted in red.

Can you guess the word within six tries?

## Installation Instructions

To run the Wordle game, you'll need to have Go (Golang) installed on your machine. Follow these steps:

1. **Install Go**:
    - Download and install Go from the [official website](https://golang.org/dl/).
    - Follow the installation instructions for your operating system.

2. **Clone the Repository**:
   ```bash
   git clone git@github.com:rafaelcx/wordle-go-app.git
   cd wordle-go-cli

3. **Run the Game:**
   ```bash
   go run cmd/main.go

## How to Play

1. **Start the Game**: Run the game using the command mentioned above.
2. **Make a Guess**: Type your five-letter word guess and press Enter.
3. **Receive Feedback**:
    - Green letters indicate correct letters in the right position.
    - Yellow letters indicate correct letters in the wrong position.
    - Red letters indicate letters not in the word.
4. **Continue Guessing**: Use the feedback to inform your next guess. You have a total of six attempts to find the correct word.
5. **Win or Lose**: If you guess the word within six attempts, you win! If not, the correct word will be revealed at the end.

Have fun and enjoy playing Wordle!
