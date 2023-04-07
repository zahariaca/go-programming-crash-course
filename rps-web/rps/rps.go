package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

var drawMessages = []string{
	"Uh oh. Try again",
	"Great mind think alike.",
	"Nobody wins but you can try again.",
}

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	var message string

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessages[rand.Intn(3)]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = "You should buy a lottery ticket!"
	} else {
		roundResult = "Computer wins!"
		message = "Too bad!"
	}

	return Round{
		Message:        message,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
}
