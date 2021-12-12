package day10

import (
	"fmt"
	"sort"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 10")
	example := d.ReadLines("./days/day10/example")
	data := d.ReadLines("./days/day10/input")

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
	points := 0

	for i := range lines {
		if verbose {
			fmt.Println("* Evaluating " + lines[i])
		}

		chars := strings.Split(lines[i], "")

		stack := make([]string, 0)

	line_loop:
		for j := range chars {
			switch chars[j] {
			case "(":
				fallthrough
			case "[":
				fallthrough
			case "{":
				fallthrough
			case "<":
				stack = append([]string{chars[j]}, stack...)
			case ")":
				if stack[0] != "(" {
					if verbose {
						fmt.Println("Syntax error (unexpected )): 3 points")
					}
					points += 3
					break line_loop
				}

				stack = stack[1:]
			case "]":
				if stack[0] != "[" {
					if verbose {
						fmt.Println("Syntax error (unexpected ]): 57 points")
					}
					points += 57
					break line_loop
				}

				stack = stack[1:]
			case "}":
				if stack[0] != "{" {
					if verbose {
						fmt.Println("Syntax error (unexpected }): 1197 points")
					}
					points += 1197
					break line_loop
				}

				stack = stack[1:]
			case ">":
				if stack[0] != "<" {
					if verbose {
						fmt.Println("Syntax error (unexpected >): 25137 points")
					}
					points += 25137
					break line_loop
				}

				stack = stack[1:]
			}
		}
	}

	if verbose {
		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("Result: Total syntax errors score is %d", points))
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	var totals []int

	for i := range lines {
		if verbose {
			fmt.Println("* Evaluating " + lines[i])
		}

		invalid := false

		chars := strings.Split(lines[i], "")

		stack := make([]string, 0)

	line_loop:
		for j := range chars {
			switch chars[j] {
			case "(":
				fallthrough
			case "[":
				fallthrough
			case "{":
				fallthrough
			case "<":
				stack = append([]string{chars[j]}, stack...)
			case ")":
				if stack[0] != "(" {
					if verbose {
						fmt.Println("Syntax error (unexpected )): ignoring")
					}
					invalid = true
					break line_loop
				}

				stack = stack[1:]
			case "]":
				if stack[0] != "[" {
					if verbose {
						fmt.Println("Syntax error (unexpected ]): ignoring")
					}
					invalid = true
					break line_loop
				}

				stack = stack[1:]
			case "}":
				if stack[0] != "{" {
					if verbose {
						fmt.Println("Syntax error (unexpected }): ignoring")
					}
					invalid = true
					break line_loop
				}

				stack = stack[1:]
			case ">":
				if stack[0] != "<" {
					if verbose {
						fmt.Println("Syntax error (unexpected >): ignoring")
					}
					invalid = true
					break line_loop
				}

				stack = stack[1:]
			}
		}

		if !invalid {
			if verbose {
				fmt.Println(fmt.Sprintf("Missing stack : %s", strings.Join(stack, ",")))
			}

			points := 0
			for k := range stack {
				switch stack[k] {
				case "(":
					points = points*5 + 1
				case "[":
					points = points*5 + 2
				case "{":
					points = points*5 + 3
				case "<":
					points = points*5 + 4
				}
			}

			if verbose {
				fmt.Println(fmt.Sprintf("Partial score: %d", points))
			}

			totals = append(totals, points)
		}
	}

	if verbose {
		fmt.Println()
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] < totals[j]
	})

	fmt.Println(fmt.Sprintf("Result: Middle autocomplete score is %d", totals[int(float64(len(totals))/2.0-0.5)]))
	fmt.Println()
}
