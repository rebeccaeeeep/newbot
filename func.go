package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yanzay/tbot/v2"
)

var picks = []string{"rock", "paper", "scissors", "lizard", "spock"} // choices from where the bot picks

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func makeButtons() *tbot.InlineKeyboardMarkup {
	// Create butttons with visible Text and CallbackData as a value.
	btnRock := tbot.InlineKeyboardButton{
		Text:         "Rock 👊🏻",
		CallbackData: "rock",
	}
	btnPaper := tbot.InlineKeyboardButton{
		Text:         "Paper ✋🏻",
		CallbackData: "paper",
	}
	btnScissors := tbot.InlineKeyboardButton{
		Text:         "Scissors✌🏻",
		CallbackData: "scissors",
	}
	btnLizard := tbot.InlineKeyboardButton{
		Text:         "Lizard 👌🏻",
		CallbackData: "lizard",
	}
	btnSpock := tbot.InlineKeyboardButton{
		Text:         "Spock 🖖🏻",
		CallbackData: "spock",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnRock, btnPaper, btnScissors, btnLizard, btnSpock},
		},
	}
}

func (a *application) draw(humanMove string) (msg string) {
	var result string
	botMove := picks[rand.Intn(len(picks))] // Generaate a random option for the bot

	// Determine the outcome
	switch humanMove {
	case botMove:
		result = "drew"
		a.draws++
	case options[botMove]:
		result = "lost"
		a.losses++
	case options2[botMove]:
		result = "lost"
	default:
		result = "won"
		a.wins++
	}
	msg = fmt.Sprintf("You %s! You chose %s and I chose %s.", result, humanMove, botMove)
	return
}
