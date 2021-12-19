package day15

import "fmt"

type Node struct {
	x    int
	y    int
	risk int
	cost int
}

func (n *Node) ToString() string {
	return fmt.Sprintf("%d", n.x) + "," + fmt.Sprintf("%d", n.y)
}

func (n *Node) EstimateTo(targetX int, targetY int) int {
	return (targetX - n.x) + (targetY - n.y)
}
