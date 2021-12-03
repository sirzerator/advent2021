package day1

import (
	"fmt"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 1")
	example := d.ReadLines("./days/day1/example")
	data := d.ReadLines("./days/day1/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(d.ArrayToInteger(example), verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(d.ArrayToInteger(example), verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(d.ArrayToInteger(data), verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(d.ArrayToInteger(data), verbose)
}

func executePart1(values []int, verbose bool) {
	var (
		increments int
		lastValue  int
	)
	increments = 0
	lastValue = values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > lastValue {
			increments += 1

			if verbose {
				fmt.Println(fmt.Sprintf("%d is larger than %d: increments raised to %d",
					values[i], lastValue, increments))
			}
		} else if verbose {
			fmt.Println(fmt.Sprintf("%d is smaller or equal to %d", values[i], lastValue))
		}

		lastValue = values[i]
	}

	fmt.Println(fmt.Sprintf("Result: %d", increments))
	fmt.Println()
}

func executePart2(values []int, verbose bool) {
	var (
		currentWindow int
		increments    int
		lastWindow    int
	)
	increments = 0
	lastWindow = values[0] + values[1] + values[2]

	for i := 3; i < len(values); i++ {
		currentWindow = values[i-2] + values[i-1] + values[i]
		if currentWindow > lastWindow {
			increments += 1

			if verbose {
				fmt.Println(fmt.Sprintf("%d is larger than %d: increments raised to %d",
					currentWindow, lastWindow, increments))
			}
		} else if verbose {
			fmt.Println(fmt.Sprintf("%d is smaller or equal to %d", currentWindow, lastWindow))
		}

		lastWindow = currentWindow
	}

	fmt.Println(fmt.Sprintf("Result: %d", increments))
	fmt.Println()
}
