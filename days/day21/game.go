package day21

import "fmt"

type Game struct {
	size            int
	players         []*Player
	dice            Dice
	diceRolls       int
	nextPlayerIndex int
}

func NewGame(size int, players []*Player, dice Dice) *Game {
	game := new(Game)

	game.size = size
	game.players = players
	game.dice = dice
	game.nextPlayerIndex = 0

	return game
}

func (g *Game) nextTurn() {
	player := g.players[g.nextPlayerIndex]

	roll1 := g.dice.throw()
	roll2 := g.dice.throw()
	roll3 := g.dice.throw()
	g.diceRolls += 3

	rolls := roll1 + roll2 + roll3

	player.position = (player.position + rolls) % g.size
	if player.position == 0 {
		player.position = g.size
	}

	player.score += player.position

	g.nextPlayerIndex = g.nextPlayerIndex + 1
	if g.nextPlayerIndex >= len(g.players) {
		g.nextPlayerIndex = 0
	}
}

func (g Game) ToString() string {
	out := ""

	for i := 0; i < len(g.players); i++ {
		out += fmt.Sprintf(
			"Player %d:\n  - Score: %d\n  - Position: %d\n",
			(i + 1),
			g.players[i].score,
			g.players[i].position,
		)
	}

	return out
}

func (g Game) IsDone() bool {
	for i := 0; i < len(g.players); i++ {
		if g.players[i].score >= 1000 {
			return true
		}
	}

	return false
}
