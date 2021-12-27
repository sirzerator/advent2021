package day20

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 20")
	example := d.ReadLines("./days/day20/example")
	data := d.ReadLines("./days/day20/input")

	d.PrintSubtitle("Example Part 1")
	execute(example, verbose)

	d.PrintSubtitle("Input Part 1")
	execute(data, verbose)
}

func execute(lines []string, verbose bool) {
	algorithm := strings.Split(lines[0], "")

	fmt.Println(algorithm)

	image := make([][]string, len(lines)-2)

	for i, line := range lines[2:] {
		chars := strings.Split(line, "")
		image[i] = chars
	}

	for i := 0; i < 50; i++ {
		if verbose {
			fmt.Println("Round ", i)
			printImage(image)
			fmt.Println("Has", countChars(image, "#"), "lit pixels")
		}

		bgChar := "."
		if algorithm[0] != "." {
			if i%2 == 1 {
				bgChar = algorithm[0]
			} else {
				bgChar = algorithm[511]
			}
		}

		image = enhanceImage(image, algorithm, bgChar, verbose)
		fmt.Println()
	}

	if verbose {
		fmt.Println("Round 50")
		printImage(image)
	}

	fmt.Println("Result:", countChars(image, "#"), "lit pixels")
	fmt.Println()
}

func printImage(image [][]string) {
	for _, line := range image {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func enhanceImage(image [][]string, algorithm []string, bgChar string, verbose bool) [][]string {
	newImage := make([][]string, len(image)+2)

	for y := -1; y < len(image)+1; y++ {
		newImage[y+1] = make([]string, len(image[0])+2)

		for x := -1; x < len(image[0])+1; x++ {
			pixelBinary := getPixelBinary(x, y, image, bgChar, false)
			newImage[y+1][x+1] = algorithm[pixelBinary]
		}
	}

	return newImage
}

func getPixelBinary(x int, y int, source [][]string, bgChar string, verbose bool) int {
	value := 0

	if verbose {
		fmt.Println("Determining value around", x, y)
		fmt.Print(getAt(source, x-1, y-1, bgChar))
	}
	if getAt(source, x-1, y-1, bgChar) == "#" {
		value += 256
	}

	if verbose {
		fmt.Print(getAt(source, x, y-1, bgChar))
	}
	if getAt(source, x, y-1, bgChar) == "#" {
		value += 128
	}

	if verbose {
		fmt.Print(getAt(source, x+1, y-1, bgChar))
		fmt.Println()
	}
	if getAt(source, x+1, y-1, bgChar) == "#" {
		value += 64
	}

	if verbose {
		fmt.Print(getAt(source, x-1, y, bgChar))
	}
	if getAt(source, x-1, y, bgChar) == "#" {
		value += 32
	}

	if verbose {
		fmt.Print(getAt(source, x, y, bgChar))
	}
	if getAt(source, x, y, bgChar) == "#" {
		value += 16
	}

	if verbose {
		fmt.Print(getAt(source, x+1, y, bgChar))
		fmt.Println()
	}
	if getAt(source, x+1, y, bgChar) == "#" {
		value += 8
	}

	if verbose {
		fmt.Print(getAt(source, x-1, y+1, bgChar))
	}
	if getAt(source, x-1, y+1, bgChar) == "#" {
		value += 4
	}

	if verbose {
		fmt.Print(getAt(source, x, y+1, bgChar))
	}
	if getAt(source, x, y+1, bgChar) == "#" {
		value += 2
	}

	if verbose {
		fmt.Print(getAt(source, x+1, y+1, bgChar))
		fmt.Println()
	}
	if getAt(source, x+1, y+1, bgChar) == "#" {
		value += 1
	}

	if verbose {
		fmt.Println(value)
		fmt.Println()
	}

	return value
}

func getAt(source [][]string, x int, y int, bgChar string) string {
	if x < 0 {
		return bgChar
	}

	if y < 0 {
		return bgChar
	}

	if y > len(source)-1 {
		return bgChar
	}

	if x > len(source[0])-1 {
		return bgChar
	}

	return source[y][x]
}

func countChars(image [][]string, char string) int {
	count := 0

	for _, line := range image {
		for _, c := range line {
			if c == char {
				count++
			}
		}
	}

	return count
}
