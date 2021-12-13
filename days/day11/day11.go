package day11

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 11")
	example := d.ReadLines("./days/day11/example")
	data := d.ReadLines("./days/day11/input")

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
	var octopi [][]int

	octopi = make([][]int, len(lines))
	for i := range lines {
		octopi[i] = d.ArrayToInteger(strings.Split(lines[i], ""))
	}

	if verbose {
		d.PrintMatrix(octopi)
	}

	flashes := 0

	for i := 1; i <= 100; i++ {
		flashes += RunGeneration(octopi)

		if verbose {
			fmt.Println(fmt.Sprintf("Step %d", i))
			d.PrintMatrix(octopi)
			fmt.Println(fmt.Sprintf("%d flashes", flashes))
			fmt.Println()
		}
	}

	fmt.Println(fmt.Sprintf("Result: %d flashes", flashes))
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	var octopi [][]int

	octopi = make([][]int, len(lines))
	for i := range lines {
		octopi[i] = d.ArrayToInteger(strings.Split(lines[i], ""))
	}

	if verbose {
		d.PrintMatrix(octopi)
	}

	generation := 0
	for {
		RunGeneration(octopi)
		generation += 1

		if verbose {
			fmt.Println(fmt.Sprintf("Generation %d", generation))
			d.PrintMatrix(octopi)
			fmt.Println()
		}

		synchronized := true
	validation:
		for i := range octopi {
			for j := range octopi[i] {
				if octopi[i][j] != 0 {
					synchronized = false
					break validation
				}
			}
		}

		if synchronized {
			break
		}
	}

	fmt.Println(fmt.Sprintf("Result: synchronized after %d generations", generation))
	fmt.Println()
}

func RunGeneration(octopi [][]int) int {
	flashes := 0
	newFlashes := -1

	d.MapMatrix(octopi, func(i int, x int, y int) int {
		return i + 1
	})

	for newFlashes != 0 {
		newFlashes = 0

		d.MapMatrix(octopi, func(i int, x int, y int) int {
			if i > 9 {
				newFlashes += 1

				if y > 0 {
					// Up
					if octopi[y-1][x] != 0 {
						octopi[y-1][x] += 1
					}
				}

				if y+1 < len(octopi) {
					// Down
					if octopi[y+1][x] != 0 {
						octopi[y+1][x] += 1
					}
				}

				if x > 0 {
					if y > 0 {
						// Up-Left
						if octopi[y-1][x-1] != 0 {
							octopi[y-1][x-1] += 1
						}
					}

					if y+1 < len(octopi) {
						// Down-Left
						if octopi[y+1][x-1] != 0 {
							octopi[y+1][x-1] += 1
						}
					}

					// Left
					if octopi[y][x-1] != 0 {
						octopi[y][x-1] += 1
					}
				}

				if x+1 < len(octopi[y]) {
					if y > 0 {
						// Up-Right
						if octopi[y-1][x+1] != 0 {
							octopi[y-1][x+1] += 1
						}
					}

					if y+1 < len(octopi) {
						// Down-Right
						if octopi[y+1][x+1] != 0 {
							octopi[y+1][x+1] += 1
						}
					}

					// Right
					if octopi[y][x+1] != 0 {
						octopi[y][x+1] += 1
					}
				}

				return 0
			}

			return i
		})

		flashes += newFlashes
	}

	return flashes
}
