package day3

import (
	"fmt"
	"strconv"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 3")
	example := d.ReadLines("./days/day3/example")
	data := d.ReadLines("./days/day3/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(example, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

func executePart1(values []string, verbose bool) {
	sums := make([]float32, len(values[0]))
	var (
		gamma      string
		epsilon    string
		gammaInt   int64
		epsilonInt int64
	)

	// Summing all columns
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			if values[i][j] == '1' {
				sums[j] += 1
			}
		}
	}

	// Building gamma and epsilon values
	for i := 0; i < len(sums); i++ {
		if (sums[i] / float32(len(values))) >= 0.5 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	// Converting from binary to decimal and multiplying
	gammaInt, _ = strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ = strconv.ParseInt(epsilon, 2, 64)

	fmt.Println("Result: " + strconv.FormatInt(gammaInt*epsilonInt, 10))
	fmt.Println()
}

func executePart2(values []string, verbose bool) {
	var (
		remainingMostCommons  []string
		remainingLeastCommons []string
		remainingMostCommon   string
		remainingLeastCommon  string
	)

	remainingMostCommons = make([]string, len(values))
	copy(remainingMostCommons, values)
	remainingLeastCommons = make([]string, len(values))
	copy(remainingLeastCommons, values)

	if verbose {
		fmt.Println("Starting with : " + strings.Join(remainingMostCommons, ", "))
	}
	for i := 0; i < len(values[0]); i++ {
		sums := make([]float32, len(values[0]))

		for i := 0; i < len(remainingMostCommons); i++ {
			for j := 0; j < len(remainingMostCommons[i]); j++ {
				if remainingMostCommons[i][j] == '1' {
					sums[j] += 1
				}
			}
		}

		n := 0
		if (sums[i] / float32(len(remainingMostCommons))) >= 0.5 {
			if verbose {
				fmt.Println(fmt.Sprintf("Most common at position %d is 1", i+1))
			}

			for _, val := range remainingMostCommons {
				if val[i] == '1' {
					remainingMostCommons[n] = val
					n++
				}
			}
		} else {
			if verbose {
				fmt.Println(fmt.Sprintf("Most common at position %d is 0", i+1))
			}

			for _, val := range remainingMostCommons {
				if val[i] == '0' {
					remainingMostCommons[n] = val
					n++
				}
			}
		}

		remainingMostCommons = remainingMostCommons[:n]

		if verbose {
			fmt.Println("Kept : " + strings.Join(remainingMostCommons, ", "))
		}

		if len(remainingMostCommons) == 1 {
			remainingMostCommon = remainingMostCommons[0]
			break
		}
	}

	remainingLeastCommons = values
	if verbose {
		fmt.Println()
		fmt.Println("Starting with : " + strings.Join(remainingLeastCommons, ", "))
	}
	for i := 0; i < len(values[0]); i++ {
		sums := make([]float32, len(values[0]))

		for i := 0; i < len(remainingLeastCommons); i++ {
			for j := 0; j < len(remainingLeastCommons[i]); j++ {
				if remainingLeastCommons[i][j] == '1' {
					sums[j] += 1
				}
			}
		}

		m := 0

		if (sums[i] / float32(len(remainingLeastCommons))) < 0.5 {
			if verbose {
				fmt.Println(fmt.Sprintf("Least common at position %d is 1", i+1))
			}

			for _, val := range remainingLeastCommons {
				if val[i] == '1' {
					remainingLeastCommons[m] = val
					m++
				}
			}
		} else {
			if verbose {
				fmt.Println(fmt.Sprintf("Least common at position %d is 0", i+1))
			}

			for _, val := range remainingLeastCommons {
				if val[i] == '0' {
					remainingLeastCommons[m] = val
					m++
				}
			}
		}

		remainingLeastCommons = remainingLeastCommons[:m]

		if verbose {
			fmt.Println("Kept : " + strings.Join(remainingLeastCommons, ", "))
		}

		if len(remainingLeastCommons) == 1 {
			remainingLeastCommon = remainingLeastCommons[0]
			break
		}
	}

	// Converting from binary to decimal and multiplying
	var (
		o2  int64
		co2 int64
	)
	o2, _ = strconv.ParseInt(remainingMostCommon, 2, 64)
	co2, _ = strconv.ParseInt(remainingLeastCommon, 2, 64)

	fmt.Println("Result: " + strconv.FormatInt(o2*co2, 10))
	fmt.Println()
}
