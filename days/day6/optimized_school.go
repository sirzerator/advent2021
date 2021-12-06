package day6

import (
	"strconv"
	"strings"
)

type OptimizedSchool struct {
	ages [9]uint
}

func NewOptimizedSchool(ages []int) *OptimizedSchool {
	s := new(OptimizedSchool)
	s.ages = [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range ages {
		s.ages[uint8(ages[i])] += 1
	}
	return s
}

func (s *OptimizedSchool) ToString() string {
	var ages []string
	for i := range s.ages {
		ages = append(ages, "Age "+strconv.Itoa(i)+": "+strconv.Itoa(int(s.ages[i])))
	}
	return strings.Join(ages, ",")
}

func (s *OptimizedSchool) Age() {
	newAges := [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0}

	newAges[0] = s.ages[1]
	newAges[1] = s.ages[2]
	newAges[2] = s.ages[3]
	newAges[3] = s.ages[4]
	newAges[4] = s.ages[5]
	newAges[5] = s.ages[6]
	newAges[6] = s.ages[7] + s.ages[0]
	newAges[7] = s.ages[8]
	newAges[8] = s.ages[0]

	s.ages = newAges
}

func (s *OptimizedSchool) Population() uint {
	population := uint(0)

	for i := range s.ages {
		population += s.ages[i]
	}

	return population
}
