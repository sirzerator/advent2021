package day9

import (
	"fmt"
	"sort"

	d "github.com/sirzerator/advent2021/days"
)

func Run(verbose bool) {
	d.PrintTitle("Running day 9")
	example := d.ReadLines("./days/day9/example")
	data := d.ReadLines("./days/day9/input")

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
	heightMap := NewHeightMap(lines)

	if verbose {
		fmt.Println(heightMap.ToString())
	}

	lowPoints := heightMap.FindLowPoints(verbose)

	sum := 0
	for i := range lowPoints {
		sum += 1 + lowPoints[i].point.height
	}

	fmt.Println(fmt.Sprintf("Result: %d low points with total value %d", len(lowPoints), sum))
	fmt.Println()
}

func executePart2(lines []string, verbose bool) {
	heightMap := NewHeightMap(lines)

	if verbose {
		fmt.Println(heightMap.ToString())
	}

	heightMap.FindLowPoints(verbose)

	heightMap.EvaluateBasinSizes(verbose)

	sort.Slice(heightMap.lowPoints, func(i, j int) bool {
		return len(heightMap.lowPoints[i].basin) > len(heightMap.lowPoints[j].basin)
	})

	mult := 1
	for i := range heightMap.lowPoints[0:3] {
		lp := heightMap.lowPoints[i]
		if verbose {
			fmt.Println(fmt.Sprintf("Basin of %d, %d is %d items large", lp.point.x, lp.point.y, len(lp.basin)))
		}
		mult *= len(lp.basin)
	}

	fmt.Println(fmt.Sprintf("Result: Product of three largest basins's size is %d", mult))
	fmt.Println()
}
