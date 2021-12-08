package day7

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 7")
	example := d.ReadLines("./days/day7/example")
	data := d.ReadLines("./days/day7/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(d.ArrayToInteger(strings.Split(example[0], ",")), verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(d.ArrayToInteger(strings.Split(example[0], ",")), verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(d.ArrayToInteger(strings.Split(data[0], ",")), verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(d.ArrayToInteger(strings.Split(data[0], ",")), verbose)
}

func executePart1(positions []int, verbose bool) {
	lastDistance := 10000000000
	position := 0

	for {
		distance := computeDistance(positions, position)

		if verbose {
			fmt.Println(fmt.Sprintf("Testing index %d : distance is %d", position, distance))
		}

		if distance < lastDistance {
			lastDistance = distance
			position += 1
			continue
		}
		break
	}

	fmt.Println(fmt.Sprintf("Result : optimal position is %d with %d fuel", position-1, lastDistance))
	fmt.Println()
}

func executePart2(positions []int, verbose bool) {
	lastDistance := 10000000000
	position := 0

	for {
		distance := computeDistance2(positions, position, verbose)

		if verbose {
			fmt.Println(fmt.Sprintf("Testing index %d : distance is %d", position, distance))
		}

		if distance < lastDistance {
			lastDistance = distance
			position += 1
			continue
		}
		break
	}

	fmt.Println(fmt.Sprintf("Result : optimal position is %d with %d fuel", position-1, lastDistance))
	fmt.Println()
}

func computeDistance(positions []int, target int) int {
	distance := 0

	for i := range positions {
		distance += simpleMax(positions[i], target) - simpleMin(positions[i], target)
	}

	return distance
}

func computeDistance2(positions []int, target int, verbose bool) int {
	totalCost := 0

	for i := range positions {
		max := simpleMax(positions[i], target)
		min := simpleMin(positions[i], target)
		distance := max - min
		cost := distance * (distance + 1) / 2

		if verbose {
			fmt.Println(fmt.Sprintf("Moving from %d to %d (%d) costs %d", positions[i], target, distance, cost))
		}

		totalCost += cost
	}

	return totalCost
}

func simpleMax(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func simpleMin(x int, y int) int {
	if x < y {
		return x
	}

	return y
}
