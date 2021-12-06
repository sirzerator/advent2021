package day6

import (
	"strconv"
	"strings"
)

type School struct {
	fishes []Fish
}

func (s *School) ToString() string {
	var ages []string
	for i := range s.fishes {
		ages = append(ages, strconv.Itoa(int(s.fishes[i].age)))
	}
	return strings.Join(ages, ",")
}

func (s *School) Age() {
	for i := range s.fishes {
		if s.fishes[i].Age() {
			s.fishes = append(s.fishes, Fish{8})
		}
	}
}

func NewSchool(ages []int) *School {
	s := new(School)
	for i := range ages {
		s.fishes = append(s.fishes, Fish{uint8(ages[i])})
	}
	return s
}
