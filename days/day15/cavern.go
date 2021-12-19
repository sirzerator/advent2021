package day15

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Cavern struct {
	nodes [][]Node
}

func NewCavern(lines []string) *Cavern {
	c := new(Cavern)

	c.nodes = make([][]Node, len(lines))

	for i, line := range lines {
		c.nodes[i] = make([]Node, len(lines[i]))

		letters := strings.Split(line, "")
		for j, risk := range letters {
			value, _ := strconv.Atoi(string(risk))
			c.nodes[i][j] = Node{j, i, value, value}
		}

	}

	return c
}

func (c *Cavern) PrintPath(path Path) string {
	output := ""
	return output
}

func (c *Cavern) ToString() string {
	output := ""
	for _, line := range c.nodes {
		for _, node := range line {
			output += strconv.Itoa(node.risk)
		}
		output += "\n"
	}
	return output
}

func (c *Cavern) FindShortestPath(targetX int, targetY int, verbose bool) (Path, error) {
	toExplore := make(Path, 0, 10)
	toExplore = append(toExplore, c.nodes[0][0])

	ancestors := make(map[string]Node)

	bestCost := make(map[string]float64)
	bestCost[toExplore[0].ToString()] = 0

	for len(toExplore) > 0 {
		node := toExplore[0]

		if node.x == targetX && node.y == targetY {
			if verbose {
				fmt.Println("Found goal, rebuilding path")
			}

			return BuildPath(ancestors, node), nil
		}

		toExplore = toExplore[1:]

		if verbose {
			fmt.Println(fmt.Sprintf("Exploring (%d, %d)", node.x, node.y))
		}

		neighbors := c.GetNeighbors(node.x, node.y)
		if verbose {
			fmt.Println(" Neighbors:")
			for _, n := range neighbors {
				fmt.Println("  " + n.ToString())
			}
		}

		for _, neighbor := range neighbors {
			_, foundCost := bestCost[neighbor.ToString()]
			if !foundCost {
				bestCost[neighbor.ToString()] = math.Inf(1)
			}

			pathCost := bestCost[node.ToString()] + float64(neighbor.cost)

			if verbose {
				fmt.Println(fmt.Sprintf(" Evaluating (%s): %d", neighbor.ToString(), int(pathCost)))
			}

			if pathCost < bestCost[neighbor.ToString()] {
				if verbose {
					fmt.Println(fmt.Sprintf("  Better than previous %f", bestCost[neighbor.ToString()]))
				}

				ancestors[neighbor.ToString()] = node
				bestCost[neighbor.ToString()] = pathCost

				if !toExplore.Contains(&neighbor) {
					if verbose {
						fmt.Println(fmt.Sprintf("  Adding (%s) to the exploration list", neighbor.ToString()))
					}
					toExplore = append(toExplore, neighbor)
				}
			}
		}
	}

	return nil, errors.New("No path found")
}

func (c *Cavern) GetNeighbors(x int, y int) []Node {
	nodes := make([]Node, 0)

	if x > 0 {
		nodes = append(nodes, c.nodes[y][x-1])
	}

	if y > 0 {
		nodes = append(nodes, c.nodes[y-1][x])
	}

	if y < len(c.nodes)-1 {
		nodes = append(nodes, c.nodes[y+1][x])
	}

	if x < len(c.nodes[y])-1 {
		nodes = append(nodes, c.nodes[y][x+1])
	}

	return nodes
}

func BuildPath(ancestors map[string]Node, node Node) Path {
	p := make(Path, 0, node.x+node.y)

	p = append(p, node)

	var (
		ancestor Node
		present  bool
	)
	for {
		if node.x == 0 && node.y == 0 {
			break
		}

		ancestor, present = ancestors[node.ToString()]

		if !present {
			break
		}

		p = append(p, ancestor)
		node = ancestor
	}

	return p
}
