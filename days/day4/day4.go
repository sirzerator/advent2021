package day4

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 4")
	example := d.ReadLines("./days/day4/example")
	data := d.ReadLines("./days/day4/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(example, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

func executePart1(rows []string, verbose bool) {
	var (
		boards []*Bingo
		calls  []int
	)
	calls, boards = prepareDay4(rows, verbose)

play:
	for c := 0; c < len(calls); c++ {
		for i := 0; i < len(boards); i++ {
			if boards[i].MarkAndTestIsWinner(calls[c]) {
				fmt.Println(fmt.Sprintf(
					"Result: board %d will win with score %d",
					i+1,
					boards[i].EvaluateScore(),
				))
				fmt.Println()
				break play
			}
		}
	}
}

func executePart2(rows []string, verbose bool) {
	var (
		boards []*Bingo
		calls  []int
		won    int
	)
	calls, boards = prepareDay4(rows, verbose)
	won = 0

play:
	for c := 0; c < len(calls); c++ {
		for i := 0; i < len(boards); i++ {
			if boards[i].winner {
				continue
			}

			if boards[i].MarkAndTestIsWinner(calls[c]) {
				won += 1
				if won == len(boards) {
					fmt.Println(fmt.Sprintf(
						"Result: board %d will win last with score %d",
						i+1,
						boards[i].EvaluateScore(),
					))
					fmt.Println()
					break play
				}
			}
		}
	}
}

func prepareDay4(rows []string, verbose bool) ([]int, []*Bingo) {
	var (
		boards   []*Bingo
		callsRow string
		rest     []string
		calls    []int
	)
	callsRow, rest = rows[0], rows[2:]
	calls = d.ArrayToInteger(strings.Split(callsRow, ","))

	for {
		boards = append(boards, NewBingo(rest[0:5]))

		if len(rest) <= 5 {
			break
		}

		rest = rest[6:]
	}

	if verbose {
		for i := 0; i < len(boards); i++ {
			fmt.Println(boards[i].ToString())
			fmt.Println()
		}
	}

	return calls, boards
}
