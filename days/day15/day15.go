package day15

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 15")
	example := d.ReadLines("./days/day15/example")
	data := d.ReadLines("./days/day15/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(example, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

func executePart1(lines []string, verbose bool) {
	cavern := NewCavern(lines)

	if verbose {
		fmt.Println(cavern.ToString())
		fmt.Println()
	}

	path, err := cavern.FindShortestPath(len(lines)-1, len(lines[0])-1, verbose)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if verbose {
		fmt.Println(path.ToString())
		fmt.Println(cavern.PrintPath(path))
		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("Result: path has a total risk of %d", path.TotalRisk()))
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	cavern := NewCavern(enlarge(lines))

	if verbose {
		fmt.Println(cavern.ToString())
		fmt.Println()
	}

	path, err := cavern.FindShortestPath(len(cavern.nodes)-1, len(cavern.nodes[0])-1, verbose)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if verbose {
		fmt.Println(path.ToString())
		fmt.Println(cavern.PrintPath(path))
		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("Result: path has a total risk of %d", path.TotalRisk()))
	fmt.Println()
}

func enlarge(lines []string) []string {
	largerMap := make([]string, len(lines)*5)

	for yInc := 0; yInc < 5; yInc++ {
		//fmt.Println(fmt.Sprintf("yInc: %d", yInc))
		for i, line := range lines {
			largerMap[i+(yInc*len(lines))] = ""
			for inc := 0; inc < 5; inc++ {
				//fmt.Println(fmt.Sprintf("inc: %d", inc))
				for _, digit := range strings.Split(line, "") {
					//fmt.Println(fmt.Sprintf("digit: %s", digit))
					digitInt, _ := strconv.Atoi(digit)
					newDigit := digitInt + inc + yInc
					//fmt.Println(fmt.Sprintf("newDigit: %d", newDigit))
					if newDigit > 9 {
						newDigit += 1
						newDigit %= 10
					}
					largerMap[i+(yInc*(len(lines)))] += fmt.Sprintf("%d", newDigit)
					//fmt.Println(largerMap)
				}
			}
		}
	}

	return largerMap
}
