package day15

import (
	"errors"
	"fmt"
)

type Path []Node

func (p Path) Len() int {
	return len(p)
}

func (p Path) Less(i, j int) bool {
	return p[i].risk < p[j].risk
}

func (p Path) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Path) ComputeRisk() int {
	risk := 0

	for _, node := range p {
		risk += node.risk
	}

	return risk
}

func (p Path) ToString() string {
	output := ""

	for _, node := range p {
		if output != "" {
			output += " -> "
		}

		output += fmt.Sprintf("(%d, %d)", node.x, node.y)
	}

	return output
}

func (p Path) GetNodeIndex(node *Node) (int, error) {
	for i, pathNode := range p {
		if pathNode.x == node.x && pathNode.y == node.y {
			return i, nil
		}
	}

	return -1, errors.New("Not found")
}

func (p *Path) Contains(node *Node) bool {
	_, err := p.GetNodeIndex(node)

	if err != nil {
		return false
	}

	return true
}

func (p *Path) ContainsCloser(node *Node) bool {
	index, err := p.GetNodeIndex(node)

	if err != nil {
		return false
	}

	return (*p)[index].cost < node.cost
}

func (p Path) TotalRisk() int {
	sum := 0

	for _, node := range p {
		sum += node.risk
	}

	sum -= p[len(p)-1].risk

	return sum
}
