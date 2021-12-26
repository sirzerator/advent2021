package day18

import (
	"fmt"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 18")
	example1 := d.ReadLines("./days/day18/example1")
	example2 := d.ReadLines("./days/day18/example2")
	example3 := d.ReadLines("./days/day18/example3")
	example4 := d.ReadLines("./days/day18/example4")
	example5 := d.ReadLines("./days/day18/example5")
	example6 := d.ReadLines("./days/day18/example6")
	data := d.ReadLines("./days/day18/input")

	d.PrintSubtitle("Example Part 1 (I)")
	executePart1(example1, verbose)
	fmt.Println()
	d.PrintSubtitle("Example Part 1 (II)")
	executePart1(example2, verbose)
	fmt.Println()
	d.PrintSubtitle("Example Part 1 (III)")
	executePart1(example3, verbose)
	fmt.Println()
	d.PrintSubtitle("Example Part 1 (IV)")
	executePart1(example4, verbose)
	fmt.Println()
	d.PrintSubtitle("Example Part 1 (V)")
	executePart1(example5, verbose)
	fmt.Println()
	d.PrintSubtitle("Example Part 1 (VI)")
	executePart1(example6, verbose)
	fmt.Println()

	d.PrintSubtitle("Example Part 2 (VI)")
	executePart2(example6, verbose)
	fmt.Println()

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

func executePart1(lines []string, verbose bool) {
	sailfishes := make([]*Sailfish, 0, len(lines))

	for _, line := range lines {
		sailfishes = append(sailfishes, NewSailfishFromString(line, false))
	}

	for len(sailfishes) > 1 {
		var restOfSailfishes []*Sailfish
		if len(sailfishes) > 2 {
			restOfSailfishes = sailfishes[2:]
		} else {
			restOfSailfishes = []*Sailfish{}
		}

		if verbose {
			fmt.Println(sailfishes[0].ToString())
			fmt.Println("+ " + sailfishes[1].ToString())
		}

		sailfishes = append(
			[]*Sailfish{sailfishes[0].AddSailfish(sailfishes[1])},
			restOfSailfishes...,
		)

		if verbose {
			fmt.Println("= " + sailfishes[0].ToString())
			fmt.Println()
		}

		sailfishes[0].ResolveSailfish(verbose)

		if verbose {
			fmt.Println("=> " + sailfishes[0].ToString())
			fmt.Println()
		}
	}

	fmt.Println("Result: final sum magnitude is ", sailfishes[0].GetMagnitude())
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	largestMagnitude := 0
	sailfishAMemo := ""
	sailfishBMemo := ""

	for a, linea := range lines {
		for b, lineb := range lines {
			if a == b {
				continue
			}

			sailfishA := NewSailfishFromString(linea, false)
			sailfishB := NewSailfishFromString(lineb, false)

			sumSailfish := sailfishA.AddSailfish(sailfishB)
			sumSailfish.ResolveSailfish(verbose)

			magnitude := sumSailfish.GetMagnitude()

			if verbose {
				fmt.Println(fmt.Sprintf("Magnitude of %s + %s = %s => %d", linea, lineb, sumSailfish.ToString(), sumSailfish.GetMagnitude()))
				fmt.Println()
			}

			if magnitude > largestMagnitude {
				largestMagnitude = magnitude
				sailfishAMemo = linea
				sailfishBMemo = lineb
			}
		}
	}

	fmt.Println("Largest magnitude is ", sailfishAMemo, " + ", sailfishBMemo, " => ", largestMagnitude)
	fmt.Println()
}
