package day9

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

type HeightMap struct {
	points [][]Point

	lowPoints []LowPoint
}

type Point struct {
	x      int
	y      int
	height int
	pMap   *HeightMap
}

type LowPoint struct {
	point Point

	basin []Point
}

func (lp *LowPoint) EvaluateBasinSize(verbose bool) []Point {
	if verbose {
		fmt.Println(fmt.Sprintf("Starting at %d, %d", lp.point.x+1, lp.point.y+1))
	}

	var (
		pointsToExplore []Point
		points          [][]Point
		basin           []Point
	)

	basin = make([]Point, 0)
	points = lp.point.pMap.points
	pointsToExplore = []Point{lp.point}

	for len(pointsToExplore) > 0 {
		np := pointsToExplore[0]
		pointsToExplore = pointsToExplore[1:]

		if explored(basin, np) {
			continue
		}

		basin = append(basin, np)

		// Left
		if np.x > 0 && points[np.y][np.x-1].height < 9 {
			point := points[np.y][np.x-1]

			if !explored(basin, point) {
				pointsToExplore = append(pointsToExplore, point)

				if verbose {
					fmt.Println(fmt.Sprintf("Left: Adding %d, %d to basin", point.x+1, point.y+1))
				}
			}
		}

		// Right
		if np.x+1 < len(points[np.y]) && points[np.y][np.x+1].height < 9 {
			point := points[np.y][np.x+1]
			if !explored(basin, point) {
				pointsToExplore = append(pointsToExplore, point)

				if verbose {
					fmt.Println(fmt.Sprintf("Right: Adding %d, %d to basin", point.x+1, point.y+1))
				}
			}
		}

		// Up
		if np.y > 0 && points[np.y-1][np.x].height < 9 {
			point := points[np.y-1][np.x]
			if !explored(basin, point) {
				pointsToExplore = append(pointsToExplore, point)

				if verbose {
					fmt.Println(fmt.Sprintf("Up: Adding %d, %d to basin", point.x+1, point.y+1))
				}
			}
		}

		// Down
		if np.y+1 < len(points) && points[np.y+1][np.x].height < 9 {
			point := points[np.y+1][np.x]
			if !explored(basin, point) {
				pointsToExplore = append(pointsToExplore, point)

				if verbose {
					fmt.Println(fmt.Sprintf("Down: Adding %d, %d to basin", point.x+1, point.y+1))
				}
			}
		}
	}

	if verbose {
		fmt.Println(fmt.Sprintf("Basin is %d elements large", len(basin)))
	}
	lp.basin = basin

	return basin
}

func NewHeightMap(lines []string) *HeightMap {
	hm := new(HeightMap)

	hm.lowPoints = make([]LowPoint, 0)

	hm.points = make([][]Point, len(lines))
	for y := range lines {
		points := d.ArrayToInteger(strings.Split(lines[y], ""))

		hm.points[y] = make([]Point, len(points))

		for x := range points {
			hm.points[y][x] = Point{x, y, points[x], hm}
		}
	}

	return hm
}

func (hm *HeightMap) FindLowPoints(verbose bool) []LowPoint {
	hm.lowPoints = make([]LowPoint, 0)

	for y := 0; y < len(hm.points); y++ {
		for x := 0; x < len(hm.points[y]); x++ {
			if verbose {
				fmt.Println(fmt.Sprintf("Evaluating %d, %d", x+1, y+1))
			}

			if hm.IsLowPoint(x, y, verbose) {
				hm.lowPoints = append(hm.lowPoints, LowPoint{hm.points[y][x], make([]Point, 0)})
			}
		}
	}

	return hm.lowPoints
}

func (hm *HeightMap) EvaluateBasinSizes(verbose bool) []LowPoint {
	for i := range hm.lowPoints {
		hm.lowPoints[i].EvaluateBasinSize(verbose)
	}

	return hm.lowPoints
}

func (hm *HeightMap) IsLowPoint(x int, y int, verbose bool) bool {
	value := hm.points[y][x].height
	if verbose {
		fmt.Println(fmt.Sprintf("Value is %d", value))
	}

	if verbose && x > 0 {
		fmt.Println(fmt.Sprintf("Left: %d", hm.points[y][x-1].height))
	}
	if x > 0 && hm.points[y][x-1].height <= value {
		if verbose {
			fmt.Println(fmt.Sprintf("Left is smaller: %d", hm.points[y][x-1].height))
		}
		return false
	}

	if verbose && x+1 < len(hm.points[y]) {
		fmt.Println(fmt.Sprintf("Right: %d", hm.points[y][x+1].height))
	}
	if x+1 < len(hm.points[y]) && hm.points[y][x+1].height <= value {
		if verbose {
			fmt.Println(fmt.Sprintf("Right is smaller: %d", hm.points[y][x+1].height))
		}
		return false
	}

	if verbose && y > 0 {
		fmt.Println(fmt.Sprintf("Up: %d", hm.points[y-1][x].height))
	}
	if y > 0 && hm.points[y-1][x].height <= value {
		if verbose {
			fmt.Println(fmt.Sprintf("Up is smaller: %d", hm.points[y-1][x].height))
		}
		return false
	}

	if verbose && y+1 < len(hm.points) {
		fmt.Println(fmt.Sprintf("Down: %d", hm.points[y+1][x].height))
	}
	if y+1 < len(hm.points) && hm.points[y+1][x].height <= value {
		if verbose {
			fmt.Println(fmt.Sprintf("Down is smaller: %d", hm.points[y+1][x].height))
		}
		return false
	}

	if verbose {
		fmt.Println("Is a LOWPOINT")
	}

	return true
}

func (hm *HeightMap) ToString() string {
	output := ""

	for i := range hm.points {
		for j := range hm.points[i] {
			output += fmt.Sprint(hm.points[i][j].height) + " "
		}
		output += "\n"
	}
	return output
}

func explored(points []Point, p Point) bool {
	for _, explored := range points {
		if explored == p {
			return true
		}
	}
	return false
}
