package day21

import (
	"strconv"
	"strings"
)

type Player struct {
	position int
	score    int
}

func NewPlayer(line string) *Player {
	player := new(Player)

	parts := strings.Split(line, " ")
	position, _ := strconv.Atoi(parts[len(parts)-1])

	player.position = position
	player.score = 0

	return player
}
