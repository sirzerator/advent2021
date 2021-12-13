package day12

import (
	"fmt"
	"strings"
)

type CaveSystem struct {
	caves    []*Cave
	cavesMap map[string]*Cave
}

func NewCaveSystem(lines []string) *CaveSystem {
	cs := new(CaveSystem)
	cs.cavesMap = make(map[string]*Cave)
	cs.caves = make([]*Cave, 0)

	for i := range lines {
		parts := strings.Split(lines[i], "-")
		left, right := parts[0], parts[1]

		var leftCave *Cave
		_, leftPresent := cs.cavesMap[left]
		if !leftPresent {
			leftCave = NewCave(left)
			cs.caves = append(cs.caves, leftCave)
			cs.cavesMap[left] = leftCave
		}
		leftCave = cs.cavesMap[left]

		var rightCave *Cave
		_, rightPresent := cs.cavesMap[right]
		if !rightPresent {
			rightCave = NewCave(right)
			cs.caves = append(cs.caves, rightCave)
			cs.cavesMap[right] = rightCave
		}
		rightCave = cs.cavesMap[right]

		if rightCave.label != "end" && leftCave.label != "start" {
			cs.cavesMap[right].reaches = append(cs.cavesMap[right].reaches, leftCave)
		}
		if leftCave.label != "end" && rightCave.label != "start" {
			cs.cavesMap[left].reaches = append(cs.cavesMap[left].reaches, rightCave)
		}
	}

	return cs
}

func (cs *CaveSystem) Explore(canVisitOneSmallCaveTwice bool, verbose bool) []string {
	paths := make([]string, 0)

	var pathSoFar []*Cave
	pathSoFar = append(pathSoFar, cs.cavesMap["start"])
	cavePaths := CaveSystemDepthFirstSearch(pathSoFar, canVisitOneSmallCaveTwice, verbose)

	for _, path := range cavePaths {
		paths = append(paths, PathToString(path))
	}

	return paths
}

func CaveSystemDepthFirstSearch(pathSoFar []*Cave, canVisitOneSmallCaveTwice bool, verbose bool) [][]*Cave {
	lastCave := pathSoFar[len(pathSoFar)-1]
	if verbose {
		fmt.Println(PathToString(pathSoFar))
	}

	if lastCave.label == "end" {
		if verbose {
			fmt.Println("** Reached end: this is a path")
			fmt.Println()
		}
		return [][]*Cave{pathSoFar}
	}

	candidates := make([]*Cave, 0)

	if canVisitOneSmallCaveTwice {
		for i := range lastCave.reaches {
			if lastCave.reaches[i].big {
				candidates = append(candidates, lastCave.reaches[i])
			} else if !PathContainsDuplicates(pathSoFar) {
				candidates = append(candidates, lastCave.reaches[i])
			} else if !PathContains(lastCave.reaches[i].label, pathSoFar) {
				candidates = append(candidates, lastCave.reaches[i])
			}
		}
	} else {
		for i := range lastCave.reaches {
			if lastCave.reaches[i].big {
				candidates = append(candidates, lastCave.reaches[i])
			} else if !PathContains(lastCave.reaches[i].label, pathSoFar) {
				candidates = append(candidates, lastCave.reaches[i])
			}
		}
	}

	var paths [][]*Cave

	if verbose {
		fmt.Print("Candidates: ")
		for _, cand := range candidates {
			fmt.Print(cand.label + " ")
		}
		fmt.Println()
	}
	for _, candidate := range candidates {
		if verbose {
			fmt.Println("-> Exploring candidate " + candidate.label)
		}
		newPaths := CaveSystemDepthFirstSearch(append(pathSoFar, candidate), canVisitOneSmallCaveTwice, verbose)
		paths = append(paths, newPaths...)
	}

	if verbose {
		for _, path := range paths {
			fmt.Println(PathToString(path))
		}
	}
	return paths
}

func (cs *CaveSystem) ToString() string {
	str := ""

	for i := range cs.caves {
		cave := cs.caves[i]
		str += cave.label + " "
		if cave.big {
			str += "(big)"
		} else {
			str += "(small)"
		}
		str += "\n"

		for j := range cave.reaches {
			str += " -> " + cave.reaches[j].label + "\n"
		}
	}

	return str
}

func PathContains(label string, path []*Cave) bool {
	for _, cave := range path {
		if cave.label == label {
			return true
		}
	}

	return false
}

func PathContainsDuplicates(path []*Cave) bool {
	var labelsMap map[string]int
	labelsMap = make(map[string]int)

	for _, cave := range path {
		if cave.big {
			continue
		}

		_, present := labelsMap[cave.label]

		if !present {
			labelsMap[cave.label] = 0
		}

		labelsMap[cave.label] += 1

		if labelsMap[cave.label] > 1 {
			return true
		}
	}

	return false
}

func PathToString(path []*Cave) string {
	str := ""
	for i, cave := range path {
		if i > 0 {
			str += " -> " + cave.label
		} else {
			str += cave.label
		}
	}
	str += "\n"
	return str
}
