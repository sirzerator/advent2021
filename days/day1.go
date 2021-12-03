package days

import (
	"fmt"
)

func day1(verbose bool) {
	PrintTitle("Running day 1")
	PrintSubtitle("Example Part 1")
	example := ReadLines("./days/day1/example")
	executePart1(ArrayToInteger(example), verbose)
	PrintSubtitle("Example Part 2")
	executePart2(ArrayToInteger(example), verbose)

	data := ReadLines("./days/day1/input")
	PrintSubtitle("Input Part 1")
	executePart1(ArrayToInteger(data), verbose)
	PrintSubtitle("Input Part 2")
	executePart2(ArrayToInteger(data), verbose)
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
