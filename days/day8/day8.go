package day8

import (
	"fmt"
	"math"
	"sort"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 8")
	example := d.ReadLines("./days/day8/example")
	shortExample := d.ReadLines("./days/day8/short_example")
	data := d.ReadLines("./days/day8/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, verbose)
	d.PrintSubtitle("Example Part 2 (short)")
	executePart2(shortExample, verbose)
	d.PrintSubtitle("Example Part 2 (long)")
	executePart2(example, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, verbose)
}

func executePart1(lines []string, verbose bool) {
	count := 0

	for i := range lines {
		parts := strings.Split(lines[i], " | ")
		right := parts[1]

		rightDigits := strings.Split(right, " ")

		for j := range rightDigits {
			if verbose {
				fmt.Println(fmt.Sprintf("Evaluating %s...", rightDigits[j]))
			}

			switch len(rightDigits[j]) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				if verbose {
					fmt.Println(fmt.Sprintf("Length is %d: counting one", len(rightDigits[j])))
				}

				count += 1
			}
		}
	}

	fmt.Println(fmt.Sprintf("Result : %d are 1, 4, 7, or 8", count))
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	total := 0

	for i := range lines {
		var (
			digit1Parts []string
			digit4Parts []string
		)

		digit0 := ""
		digit1 := ""
		digit2 := ""
		digit3 := ""
		digit4 := ""
		digit5 := ""
		digit6 := ""
		digit7 := ""
		digit8 := ""
		digit9 := ""

		digitsMap := make(map[string]int)

		parts := strings.Split(lines[i], " | ")
		digits := strings.Split(parts[0], " ")

		// First pass
		if verbose {
			fmt.Println()
			fmt.Println("First pass (easy):")
			fmt.Println()
		}

		for j := range digits {
			switch len(digits[j]) {
			case 2:
				if digit1 != "" {
					if verbose {
						fmt.Println("Already found: 1")
					}
					continue
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
					fmt.Println("Length is 2, so it's a 1")
				}

				digit1 = digits[j]
				digit1Parts = strings.Split(digit1, "")
				digitsMap[SortString(digit1)] = 1
			case 3:
				if digit7 != "" {
					if verbose {
						fmt.Println("Already found: 7")
					}
					continue
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
					fmt.Println("Length is 3, so it's a 7")
				}

				digit7 = digits[j]
				digitsMap[SortString(digit7)] = 7
			case 4:
				if digit4 != "" {
					if verbose {
						fmt.Println("Already found: 4")
					}
					continue
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
					fmt.Println("Length is 4, so it's a 4")
				}

				digit4 = digits[j]
				digit4Parts = strings.Split(digit4, "")
				digitsMap[SortString(digit4)] = 4
			case 7:
				if digit8 != "" {
					if verbose {
						fmt.Println("Already found: 8")
					}
					continue
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
					fmt.Println("Length is 7, so it's a 8")
				}

				digit8 = digits[j]
				digitsMap[SortString(digit8)] = 8
			}
		}

		// Second pass
		if verbose {
			fmt.Println()
			fmt.Println("Second pass (deduction):")
			fmt.Println()
		}

		for j := range digits {
			switch len(digits[j]) {
			case 5:
				if digit3 != "" {
					continue
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
					fmt.Println("Length is 5: either 2, 3, 5")
				}

				currentDigitParts := strings.Split(digits[j], "")

				differences := countDifferences(currentDigitParts, digit1Parts)

				if verbose {
					fmt.Println(
						fmt.Sprintf("%d differences between %s and %s", differences, digits[j], digit1),
					)
				}

				if differences == 3 {
					if verbose {
						fmt.Println("Length is 5 and 3 differences with 1, so it's a 3")
					}

					digit3 = digits[j]
					digitsMap[SortString(digit3)] = 3
				}
			case 6:
				if digit6 != "" {
					continue
				}

				if verbose {
					fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
					fmt.Println("Length is 6: either 0, 6, 9")
				}

				currentDigitParts := strings.Split(digits[j], "")

				differences := countDifferences(currentDigitParts, digit1Parts)

				if verbose {
					fmt.Println(
						fmt.Sprintf("%d differences between %s and %s", differences, digits[j], digit1),
					)
				}

				if differences == 5 {
					if verbose {
						fmt.Println("Length is 6 and 5 differences with 1, so it's a 6")
					}

					digit6 = digits[j]
					digitsMap[SortString(digit6)] = 6
				}
			}
		}

		// Third pass
		if verbose {
			fmt.Println()
			fmt.Println("Third pass (elimination):")
			fmt.Println()
		}

		for k := 0; k < 2; k++ {
			for j := range digits {
				if verbose {
				}

				switch len(digits[j]) {
				case 5:
					if digit2 != "" && digit5 != "" {
						continue
					}

					if digit3 == digits[j] {
						continue
					}

					if verbose {
						fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
						fmt.Println("Length is 5: either 2 or 5")
					}

					currentDigitParts := strings.Split(digits[j], "")

					if digit2 != "" && digit5 != "" {
						if verbose {
							fmt.Println("Already found 3 and 5, so the remaining length of 5 is 2")
						}

						digit2 = digits[j]
						digitsMap[SortString(digit2)] = 2
						continue
					}

					if digit2 != "" {
						if verbose {
							fmt.Println("Already found 3 and 2, so the remaining length of 5 is 5")
						}

						digit5 = digits[j]
						digitsMap[SortString(digit5)] = 5
						continue
					}

					differences := countDifferences(currentDigitParts, digit4Parts)

					if verbose {
						fmt.Println(
							fmt.Sprintf("%d differences between %s and %s", differences, digits[j], digit4),
						)
					}

					if differences == 2 {
						if verbose {
							fmt.Println("Length is 5 and 2 differences with 4, so it's a 5")
						}

						digit5 = digits[j]
						digitsMap[SortString(digit5)] = 5
					} else if differences == 3 {
						if verbose {
							fmt.Println("Length is 5 and 3 differences with 4, so it's a 2")
						}

						digit2 = digits[j]
						digitsMap[SortString(digit2)] = 2
					}
				case 6:
					if digit9 != "" && digit0 != "" {
						continue
					}

					if verbose {
						fmt.Println(fmt.Sprintf("Evaluating %s...", digits[j]))
						fmt.Println("Length is 6: either 0 or 9")
					}

					if digits[j] == digit6 {
						if verbose {
							fmt.Println("It's a 6")
						}
						continue
					}

					if digit9 != "" {
						if verbose {
							fmt.Println("Already found 9, so only 0 is remaining")
						}

						digit0 = digits[j]
						digitsMap[SortString(digit0)] = 0
						continue
					}

					if digit0 != "" {
						if verbose {
							fmt.Println("Already found 0, so only 9 is remaining")
						}

						digit9 = digits[j]
						digitsMap[SortString(digit9)] = 9
						continue
					}

					currentDigitParts := strings.Split(digits[j], "")

					differences := countDifferences(currentDigitParts, digit4Parts)

					if verbose {
						fmt.Println(
							fmt.Sprintf("%d differences between %s and %s", differences, digits[j], digit1),
						)
					}

					if differences == 2 {
						if verbose {
							fmt.Println("Length is 6 and 2 differences with 4, so it's a 9")
						}

						digit9 = digits[j]
						digitsMap[SortString(digit9)] = 9
					} else if differences == 3 {
						if verbose {
							fmt.Println("Length is 6 and 3 differences with 4, so it's a 0")
						}

						digit0 = digits[j]
						digitsMap[SortString(digit0)] = 0
					}
				}
			}
		}

		if verbose {
			fmt.Println()
			fmt.Println(digitsMap)
			fmt.Println("0: " + digit0)
			fmt.Println("1: " + digit1)
			fmt.Println("2: " + digit2)
			fmt.Println("3: " + digit3)
			fmt.Println("4: " + digit4)
			fmt.Println("5: " + digit5)
			fmt.Println("6: " + digit6)
			fmt.Println("7: " + digit7)
			fmt.Println("8: " + digit8)
			fmt.Println("9: " + digit9)
			fmt.Println()
		}

		rightDigits := strings.Split(parts[1], " ")

		length := len(rightDigits)
		value := 0
		for l := range rightDigits {
			if verbose {
				fmt.Println(fmt.Sprintf("%s is %d", rightDigits[l], digitsMap[SortString(rightDigits[l])]))
			}

			value += int(math.Pow(10, float64(length-l-1)) * float64(digitsMap[SortString(rightDigits[l])]))
		}

		total += value
	}

	fmt.Println(fmt.Sprintf("Result: total is %d", total))
	fmt.Println()
}

func countDifferences(a []string, b []string) int {
	common := 0

	for i := range b {
		for j := range a {
			if b[i] == a[j] {
				common += 1
			}
		}
	}

	return len(a) - common
}

func SortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
