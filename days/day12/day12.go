package day12

import (
	"fmt"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 12")
	example := d.ReadLines("./days/day12/example")
	example2 := d.ReadLines("./days/day12/example2")
	example3 := d.ReadLines("./days/day12/example3")
	data := d.ReadLines("./days/day12/input")

	d.PrintSubtitle("Example Part 1 (I)")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 1 (II)")
	executePart1(example2, verbose)
	d.PrintSubtitle("Example Part 1 (III)")
	executePart1(example3, verbose)
	d.PrintSubtitle("Example Part 2 (I)")
	executePart2(example, verbose)
	d.PrintSubtitle("Example Part 2 (II)")
	executePart2(example2, verbose)
	d.PrintSubtitle("Example Part 2 (III)")
	executePart2(example3, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

func executePart1(lines []string, verbose bool) {
	system := NewCaveSystem(lines)

	if verbose {
		fmt.Println(system.ToString())
	}

	exploration := system.Explore(false, verbose)

	if verbose {
		fmt.Println(exploration)
	}

	fmt.Println(fmt.Sprintf("Result: %d paths", len(exploration)))
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	system := NewCaveSystem(lines)

	if verbose {
		fmt.Println(system.ToString())
	}

	exploration := system.Explore(true, verbose)

	if verbose {
		fmt.Println(exploration)
	}

	fmt.Println(fmt.Sprintf("Result: %d paths", len(exploration)))
	fmt.Println()
}
