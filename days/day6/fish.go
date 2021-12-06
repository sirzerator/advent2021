package day6

type Fish struct {
	age uint8
}

// Returns boolean: has reproduced
func (f *Fish) Age() bool {
	if f.age == 0 {
		f.age = 6

		return true
	}

	f.age -= 1

	return false
}
