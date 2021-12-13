package day13

import (
	"strconv"
	"strings"
)

type Paper struct {
	lines [][]string
}

func NewPaper(width int, height int, holes []string) *Paper {
	p := new(Paper)

	p.lines = make([][]string, height)
	for i := 0; i < height; i++ {
		p.lines[i] = make([]string, width)

		for j := 0; j < width; j++ {
			p.lines[i][j] = "."
		}
	}

	for _, hole := range holes {
		parts := strings.Split(hole, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p.lines[y][x] = "#"
	}

	return p
}

func (p *Paper) FoldHorizontal(row int) {
	newLines := make([][]string, int((len(p.lines)-1)/2))

	for i := range newLines {
		newLines[i] = p.lines[i]
	}

	length := len(p.lines) - 1
	for i := length; i > row; i-- {
		for x, val := range p.lines[i] {
			if val == "#" {
				newLines[abs(i-length)][x] = "#"
			}
		}
	}

	p.lines = newLines
}

func (p *Paper) FoldVertical(column int) {
	newLines := make([][]string, len(p.lines))

	for i := range newLines {
		newLines[i] = p.lines[i][0:int((len(p.lines[i])-1)/2)]
	}

	length := len(p.lines[0]) - 1
	for y := range p.lines {
		for i := length; i > column; i-- {
			val := p.lines[y][i]
			if val == "#" {
				newLines[y][abs(i-length)] = "#"
			}
		}
	}

	p.lines = newLines
}

func (p *Paper) DotsCount() int {
	count := 0

	for _, line := range p.lines {
		for _, char := range line {
			if char == "#" {
				count += 1
			}
		}
	}

	return count
}

func (p *Paper) ToString() string {
	str := ""

	for _, line := range p.lines {
		for _, char := range line {
			str += char
		}
		str += "\n"
	}

	return str
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
