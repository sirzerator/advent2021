package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

type Point struct {
	x uint16
	y uint16
}

func Run(verbose bool) {
	d.PrintTitle("Running day 5")
	example := d.ReadLines("./days/day5/example")
	data := d.ReadLines("./days/day5/input")

	d.PrintSubtitle("Example Part 1")
	executePart1(example, 10, verbose)
	d.PrintSubtitle("Example Part 2")
	executePart2(example, 10, verbose)

	d.PrintSubtitle("Input Part 1")
	executePart1(data, 1000, verbose)
	d.PrintSubtitle("Input Part 2")
	executePart2(data, 1000, verbose)
}

func executePart1(rows []string, size uint16, verbose bool) {
	board := initializeBoard(size)

	for k := range rows {
		coordinates := extractCoordinates(rows[k], true, verbose)

		if verbose {
			fmt.Println(coordinates)
		}

		for l := 0; l < len(coordinates); l++ {
			board[coordinates[l].x][coordinates[l].y] += 1
		}

		if verbose {
			printBoard(board)
		}
	}

	fmt.Println("Result : " + strconv.Itoa(countTwosAndUp(board)) + " intersections found")
	fmt.Println()
}

func executePart2(rows []string, size uint16, verbose bool) {
	board := initializeBoard(size)

	for k := range rows {
		coordinates := extractCoordinates(rows[k], false, verbose)

		if verbose {
			fmt.Println(coordinates)
		}

		for l := 0; l < len(coordinates); l++ {
			board[coordinates[l].x][coordinates[l].y] += 1
		}

		if verbose {
			printBoard(board)
		}
	}

	fmt.Println("Result : " + strconv.Itoa(countTwosAndUp(board)) + " intersections found")
	fmt.Println()
}

func extractCoordinates(definition string, ignoreDiagonals bool, verbose bool) []Point {
	var (
		coordinates []Point
		left        string
		right       string
	)

	parts := strings.Split(definition, " -> ")
	left, right = parts[0], parts[1]

	leftParts := d.ArrayToInteger(strings.Split(left, ","))
	rightParts := d.ArrayToInteger(strings.Split(right, ","))

	if leftParts[0] == rightParts[0] { // Vertical
		if verbose {
			fmt.Println("Vertical line : " + left + " -> " + right)
		}

		interval := []uint16{uint16(leftParts[1]), uint16(rightParts[1])}
		sort.Slice(interval, func(i, j int) bool { return interval[i] < interval[j] })
		for i := interval[0]; i <= interval[1]; i++ {
			coordinates = append(coordinates, Point{uint16(leftParts[0]), uint16(i)})
		}
	} else if leftParts[1] == rightParts[1] { // Horizontal
		if verbose {
			fmt.Println("Horizontal line : " + left + " -> " + right)
		}

		interval := []uint16{uint16(leftParts[0]), uint16(rightParts[0])}
		sort.Slice(interval, func(i, j int) bool { return interval[i] < interval[j] })
		for i := interval[0]; i <= interval[1]; i++ {
			coordinates = append(coordinates, Point{uint16(i), uint16(leftParts[1])})
		}
	} else { // Diagonal
		if ignoreDiagonals {
			return coordinates
		}

		if verbose {
			fmt.Println("Diagonal line : " + left + " -> " + right)
		}

		length := Abs(leftParts[0] - rightParts[0])
		x := leftParts[0]
		y := leftParts[1]
		for i := 0; i <= length; i++ {
			coordinates = append(coordinates, Point{uint16(x), uint16(y)})

			if leftParts[0] > rightParts[0] {
				x -= 1
			} else {
				x += 1
			}

			if leftParts[1] > rightParts[1] {
				y -= 1
			} else {
				y += 1
			}
		}
	}

	return coordinates
}

func initializeBoard(size uint16) [][]uint16 {
	board := make([][]uint16, size)

	for i := range board {
		board[i] = make([]uint16, size)
		for j := range board[i] {
			board[i][j] = 0
		}
	}

	return board
}

func printBoard(board [][]uint16) {
	for i := range board {
		for j := range board[i] {
			fmt.Print(fmt.Sprintf("%4d", board[j][i]))
		}
		fmt.Println()
	}
}

func countTwosAndUp(board [][]uint16) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] >= 2 {
				count++
			}
		}
	}
	return count
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
