package day13

import (
	"fmt"
	"strconv"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 13")
	example := d.ReadLines("./days/day13/example")
	data := d.ReadLines("./days/day13/input")

	d.PrintSubtitle("Example")
	execute(example, verbose)

	d.PrintSubtitle("Input")
	execute(data, verbose)
}

func execute(lines []string, verbose bool) {
	var (
		holes        []string
		instructions [][2]string
		width        int
		height       int
	)

	holes = make([]string, 0)
	instructions = make([][2]string, 0)
	reachedInstructions := false
	width = 0
	height = 0

	for _, line := range lines {
		if line == "" {
			reachedInstructions = true
			continue
		}

		if reachedInstructions {
			parts := strings.Split(line, " ")
			instructionsParts := strings.Split(parts[2], "=")
			axis, value := instructionsParts[0], instructionsParts[1]

			if height == 0 && axis == "y" {
				height, _ = strconv.Atoi(value)
				height *= 2
				height += 1
			}
			if width == 0 && axis == "x" {
				width, _ = strconv.Atoi(value)
				width *= 2
				width += 1
			}

			instructions = append(instructions, [2]string{axis, value})
		} else {
			holes = append(holes, line)
		}
	}

	if verbose {
		fmt.Println(fmt.Sprintf("Width: %d / Height: %d", width, height))
	}

	paper := NewPaper(width, height, holes)

	for _, instruction := range instructions {
		switch instruction[0] {
		case "x":
			column, _ := strconv.Atoi(instruction[1])

			if verbose {
				fmt.Println(fmt.Sprintf("Folding vertically on column %d", column))
			}

			paper.FoldVertical(column)
		case "y":
			row, _ := strconv.Atoi(instruction[1])

			if verbose {
				fmt.Println(fmt.Sprintf("Folding horizontally on row %d", row))
			}

			paper.FoldHorizontal(row)
		}

		if verbose {
			// Writing large arrays is too slow
			if (len(paper.lines)) < 120 {
				fmt.Println(paper.ToString())
			}
			fmt.Println(fmt.Sprintf("%d dots", paper.DotsCount()))
			fmt.Println()
		}
	}

	fmt.Println("Result")
	// Writing large arrays is too slow
	if (len(paper.lines)) < 120 {
		fmt.Println(paper.ToString())
	}
	fmt.Println(fmt.Sprintf("%d dots", paper.DotsCount()))
	fmt.Println()
}
