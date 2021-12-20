package day17

import (
	"fmt"
	"strconv"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 17")
	example := d.ReadLines("./days/day17/example")
	data := d.ReadLines("./days/day17/input")

	d.PrintSubtitle("Example Part 1")
	execute(example[0], verbose)

	d.PrintSubtitle("Input Part 1")
	execute(data[0], verbose)
}

func execute(target string, verbose bool) {
	xs, ys := readLine(target)

	if verbose {
		fmt.Println("X between ", xs)
		fmt.Println("Y between ", ys)
	}

	xRange := findXRange(xs, verbose)

	if verbose {
		fmt.Println("X range: ", xRange, len(xRange))
	}

	maximumY, matchesCount := findMaximumY(xRange, ys, verbose)

	fmt.Println("Result: will reach a maximum of ", maximumY, " on a total of ", matchesCount, "possibilities")
	fmt.Println()
}

func readLine(str string) ([]int, []int) {
	parts := strings.Split(str, " ")

	xParts := strings.Split(parts[2], "=")
	xRangeParts := strings.Split(xParts[1], "..")

	x1, _ := strconv.ParseInt(xRangeParts[0], 10, 64)
	cleanX2 := strings.Replace(xRangeParts[1], ",", "", 1)
	x2, _ := strconv.ParseInt(cleanX2, 10, 64)

	yParts := strings.Split(parts[3], "=")
	yRangeParts := strings.Split(yParts[1], "..")

	y1, _ := strconv.ParseInt(yRangeParts[0], 10, 64)
	y2, _ := strconv.ParseInt(yRangeParts[1], 10, 64)

	return []int{int(x1), int(x2)}, []int{int(y1), int(y2)}
}

type Pair struct {
	value      int
	iterations int
}

func findXRange(xs []int, verbose bool) []Pair {
	tested := xs[1]
	iterations := 1
	eligible := make([]Pair, 0)

	for {
		if tested == 0 {
			break
		}

		value := computeX(tested, iterations)
		if value > xs[1] || iterations > 1000 {
			tested--
			iterations = 0
		} else if value >= xs[0] && value <= xs[1] {
			if verbose {
				fmt.Println("Eligible ", tested, iterations, value)
			}

			eligible = append([]Pair{Pair{tested, iterations}}, eligible...)
		}

		iterations++
	}

	return eligible
}

func computeX(initialVelocity int, cycles int) int {
	total := 0
	currentVelocity := initialVelocity

	for i := 1; i <= cycles; i++ {
		total = total + currentVelocity
		if currentVelocity > 0 {
			currentVelocity -= 1
		}
	}

	return total
}

func findMaximumY(validXPairs []Pair, targetYs []int, verbose bool) (int, int) {
	maximumHeight := 0
	matchesCount := 0
	seenMatches := make(map[string]bool)

	maximumYTested := 300

	for _, pair := range validXPairs {
		for i := maximumYTested; i >= targetYs[0]; i-- {
			reachingY := computeY(i, pair.iterations)

			if reachingY >= targetYs[0] && reachingY <= targetYs[1] {
				zenith := getMaximumHeight(i, pair.iterations)

				if zenith > maximumHeight {
					maximumHeight = zenith
				}

				if !seenMatches[fmt.Sprintf("%d,%d", pair.value, i)] {
					seenMatches[fmt.Sprintf("%d,%d", pair.value, i)] = true
					matchesCount++
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Reached %d with initial velocity (%d, %d)", zenith, pair.value, i))
				}
			}
		}
	}

	return maximumHeight, matchesCount
}

func computeY(velocity int, iterations int) int {
	currentVelocity := velocity
	y := 0

	for i := 0; i < iterations; i++ {
		y += currentVelocity
		currentVelocity -= 1
	}

	return y
}

func getMaximumHeight(velocity int, iterations int) int {
	currentVelocity := velocity
	y := 0

	for i := 0; i < iterations; i++ {
		y += currentVelocity
		currentVelocity -= 1

		if currentVelocity < 0 {
			return y
		}
	}

	return y
}
