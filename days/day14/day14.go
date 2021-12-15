package day14

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 14")
	example := d.ReadLines("./days/day14/example")
	data := d.ReadLines("./days/day14/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(example, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

// Naive way: simple but way too slow
func executePart1(lines []string, verbose bool) {
	str := lines[0]

	rules := make(map[string]string)

	for _, rule := range lines[2:] {
		parts := strings.Split(rule, " -> ")
		left, right := parts[0], parts[1]
		rules[left] = right
	}

	if verbose {
		fmt.Println("Template: " + str)
	}

	for i := 0; i < 10; i++ {
		initialStringParts := strings.Split(str, "")
		str = ""

		for j := 0; j < len(initialStringParts)-1; j++ {
			twoNextChars := initialStringParts[j] + initialStringParts[j+1]
			if verbose {
				//fmt.Println(fmt.Sprintf("%s -> %s", twoNextChars, rules[twoNextChars]))
			}

			str += initialStringParts[j] + rules[twoNextChars]
		}
		str += initialStringParts[len(initialStringParts)-1]

		if verbose {
			fmt.Println(fmt.Sprintf("Pass %d:", i+1))
			if verbose && len(str) < 200 {
				fmt.Println(str)
			}
		}
	}

	lfl := sortByLetterFrequency(buildLettersFrequencyMap(str))

	fmt.Println(fmt.Sprintf(
		"Result: %d (%s) - %d (%s) = %d",
		lfl[0].count,
		lfl[0].letter,
		lfl[len(lfl)-1].count,
		lfl[len(lfl)-1].letter,
		lfl[0].count-lfl[len(lfl)-1].count,
	))
	fmt.Println()
}

// Don't actually do the work: count the pairs instead of making the string
func executePart2(lines []string, verbose bool) {
	ps := NewPairStore(lines[0], lines[2:])

	if verbose {
		fmt.Println(ps.pairsCount)
	}

	for i := 0; i < 40; i++ {
		ps.ApplyRules()

		if verbose {
			fmt.Println(fmt.Sprintf("Pass %d:", i+1))
			fmt.Println(ps.pairsCount)
		}
	}

	lfl := sortByLetterFrequency(ps.CountLetters())

	if lfl[0].letter == ps.lastLetter {
		lfl[0].count++
	}

	if lfl[len(lfl)-1].letter == ps.lastLetter {
		lfl[len(lfl)-1].count++
	}

	fmt.Println(fmt.Sprintf(
		"Result: %d (%s) - %d (%s) = %d",
		lfl[0].count,
		lfl[0].letter,
		lfl[len(lfl)-1].count,
		lfl[len(lfl)-1].letter,
		lfl[0].count-lfl[len(lfl)-1].count,
	))
	fmt.Println()
}
