package day21

import (
	"fmt"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 21")
	example := d.ReadLines("./days/day21/example")
	data := d.ReadLines("./days/day21/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(example, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
}

func executePart1(lines []string, verbose bool) {
	dice := NewDeterministicDice()
	player1 := NewPlayer(lines[0])
	player2 := NewPlayer(lines[1])
	game := NewGame(10, []*Player{player1, player2}, dice)

	if verbose {
		fmt.Println("Initial state")
		fmt.Println(game.ToString())
	}

	i := 1
	for {
		if verbose {
			fmt.Println("Turn", (i+1)/2, "/ Player", game.nextPlayerIndex+1)
		}

		game.nextTurn()

		if game.IsDone() {
			break
		}

		if verbose {
			fmt.Println(game.ToString())
		}

		i++
	}

	fmt.Println(fmt.Sprintf(
		"Result: %d * %d = %d",
		game.diceRolls,
		game.players[game.nextPlayerIndex].score,
		game.diceRolls*game.players[game.nextPlayerIndex].score,
	))

	fmt.Println()
}
