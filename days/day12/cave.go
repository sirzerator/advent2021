package day12

import "regexp"

type Cave struct {
	label   string
	big     bool
	reaches []*Cave
}

func NewCave(label string) *Cave {
	c := new(Cave)
	c.label = label
	c.big, _ = regexp.MatchString(`^[A-Z]+$`, label)
	c.reaches = make([]*Cave, 0)
	return c
}
