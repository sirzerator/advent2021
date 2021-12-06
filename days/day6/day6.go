package day6

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 6")
	example := d.ReadLines("./days/day6/example")
	data := d.ReadLines("./days/day6/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(d.ArrayToInteger(strings.Split(example[0], ",")), verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(d.ArrayToInteger(strings.Split(example[0], ",")), verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(d.ArrayToInteger(strings.Split(data[0], ",")), verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(d.ArrayToInteger(strings.Split(data[0], ",")), verbose)
}

func executePart1(ages []int, verbose bool) {
	school := NewSchool(ages)

	if verbose {
		fmt.Println("Initial state: " + school.ToString())
	}

	for i := 0; i < 80; i++ {
		school.Age()
		if verbose {
			fmt.Println(fmt.Sprintf("After %2d days: %s", i+1, school.ToString()))
		}
	}

	fmt.Println(fmt.Sprintf("Result : %d fishes after 80 days", len(school.fishes)))
	fmt.Println()
}

func executePart2(ages []int, verbose bool) {
	school := NewOptimizedSchool(ages)

	if verbose {
		fmt.Println("Initial state: " + school.ToString())
	}

	for i := 0; i < 256; i++ {
		school.Age()
		if verbose {
			fmt.Println(fmt.Sprintf("After %2d days: %s", i+1, school.ToString()))
		}
	}

	fmt.Println(fmt.Sprintf("Result : %d fishes after 256 days", school.Population()))
	fmt.Println()
}
