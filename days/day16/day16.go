package day16

import (
	"fmt"
	"os"
	"strconv"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 16")
	example := d.ReadLines("./days/day16/example")
	data := d.ReadLines("./days/day16/input")

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
	sum := 0

	for _, line := range lines {
		if verbose {
			fmt.Println(line)
		}

		binary := toBitsSlice(line)

		if verbose {
			fmt.Println(binary)
		}

		packet, _ := ParseBits(binary, verbose)

		packetSum := packet.SumVersions()
		sum += packetSum

		if verbose {
			fmt.Println("Version sum: ", packetSum)
			fmt.Println()
		}
	}

	fmt.Println("Result: version sum of", sum)
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	finalValue := 0

	for _, line := range lines {
		if verbose {
			fmt.Println(line)
		}

		binary := toBitsSlice(line)

		if verbose {
			fmt.Println(binary)
		}

		packet, _ := ParseBits(binary, verbose)

		value := packet.Solve(verbose)

		if verbose {
			fmt.Println("Value: ", value)
			fmt.Println()
		}

		finalValue += value
	}

	fmt.Println("Result: ", finalValue)
	fmt.Println()
}

func toBitsSlice(hexa string) []uint64 {
	bits := []uint64{}

	for i := 0; i < len(hexa); i += 1 {
		value, err := strconv.ParseUint(hexa[i:i+1], 16, 4)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		newBits := []uint64{}
		for j := 0; j < 4; j++ {
			newBits = append([]uint64{value & 0x1}, newBits...)
			value = value >> 1
		}

		bits = append(bits, newBits...)
	}

	return bits
}

func MinOf(x int, y int) int {
	if x < y {
		return x
	}

	return y
}
