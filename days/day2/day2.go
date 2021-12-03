package day2

import (
	"fmt"
	"strconv"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 2")
	example := d.ReadLines("./days/day2/example")
	data := d.ReadLines("./days/day2/input")

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
	var (
		err        error
		offset     int
		horizontal int = 0
		vertical   int = 0
	)

	for i := 0; i < len(values); i++ {
		var str = values[i]
		if str == "" {
			continue
		}

		var parts []string = strings.Split(str, " ")
		offset, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += offset
		case "down":
			vertical += offset
		case "up":
			vertical -= offset
		}

		if verbose {
			fmt.Println("Reading " + parts[0] + " with offset " + parts[1])
			fmt.Println(fmt.Sprintf("Horizontal: %d | Vertical: %d", horizontal, vertical))
		}
	}

	fmt.Println("Result: " + strconv.Itoa(horizontal*vertical))
	fmt.Println()
}

func executePart2(values []string, verbose bool) {
	var (
		err        error
		offset     int
		horizontal int = 0
		vertical   int = 0
		aim        int = 0
	)

	for i := 0; i < len(values); i++ {
		var str = values[i]
		if str == "" {
			continue
		}

		var parts []string = strings.Split(str, " ")
		offset, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += offset
			vertical += offset * aim
		case "down":
			aim += offset
		case "up":
			aim -= offset
		}

		if verbose {
			fmt.Println("Reading " + parts[0] + " with offset " + parts[1])
			fmt.Println(fmt.Sprintf("Horizontal: %d | Vertical: %d", horizontal, vertical))
		}
	}

	fmt.Println("Result: " + strconv.Itoa(horizontal*vertical))
	fmt.Println()
}
