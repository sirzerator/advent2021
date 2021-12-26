package day18

import (
	"fmt"
	"strconv"
	"strings"
)

type Sailfish struct {
	value  int
	left   *Sailfish
	right  *Sailfish
	parent *Sailfish

	leftWithValue  *Sailfish
	rightWithValue *Sailfish
}

func NewSailfishFromString(str string, verbose bool) *Sailfish {
	var (
		previousWithValue *Sailfish
		sailfish          *Sailfish
	)

	letters := strings.Split(str, "")
	index := 0

	for index < len(letters) {
		if letters[index] == "[" {
			if sailfish == nil {
				sailfish = new(Sailfish)
				sailfish.value = -1
			} else {
				childSailfish := new(Sailfish)
				childSailfish.value = -1
				childSailfish.parent = sailfish
				sailfish = childSailfish
			}
		} else if letters[index] == "]" {
			if sailfish.parent != nil {
				if sailfish.parent.left == nil {
					sailfish.parent.left = sailfish
				} else if sailfish.parent.right == nil {
					sailfish.parent.right = sailfish
				}

				sailfish = sailfish.parent
			}
		} else if letters[index] != "," {
			if sailfish.left == nil {
				leftValue, _ := strconv.ParseInt(letters[index], 10, 64)
				sailfish.left = NewSailfish(int(leftValue))
				sailfish.left.parent = sailfish

				if previousWithValue != nil {
					previousWithValue.rightWithValue = sailfish.left
				}

				sailfish.left.leftWithValue = previousWithValue
				previousWithValue = sailfish.left
			} else {
				rightValue, _ := strconv.ParseInt(letters[index], 10, 64)
				sailfish.right = NewSailfish(int(rightValue))
				sailfish.right.parent = sailfish

				if previousWithValue != nil {
					previousWithValue.rightWithValue = sailfish.right
				}

				sailfish.right.leftWithValue = previousWithValue
				previousWithValue = sailfish.right
			}
		}

		if verbose {
			fmt.Println(index, letters[index], sailfish)
		}

		index++
	}

	return sailfish
}

func NewSailfish(value int) *Sailfish {
	s := new(Sailfish)

	s.value = value

	return s
}

func (s *Sailfish) ToString() string {
	if s.value != -1 {
		return fmt.Sprintf("%d", s.value)
	}

	return "[" + s.left.ToString() + "," + s.right.ToString() + "]"
}

func (s *Sailfish) AddSailfish(right *Sailfish) *Sailfish {
	newSailfish := new(Sailfish)
	newSailfish.value = -1

	var (
		leftmostOfRight *Sailfish
		rightmostOfLeft *Sailfish
	)

	leftmostOfRight = right.left
	for leftmostOfRight.value == -1 {
		leftmostOfRight = leftmostOfRight.left
	}

	rightmostOfLeft = s.right
	for rightmostOfLeft.value == -1 {
		rightmostOfLeft = rightmostOfLeft.right
	}

	leftmostOfRight.leftWithValue = rightmostOfLeft
	rightmostOfLeft.rightWithValue = leftmostOfRight

	newSailfish.left = s
	newSailfish.left.parent = newSailfish

	newSailfish.right = right
	newSailfish.right.parent = newSailfish

	return newSailfish
}

func (s *Sailfish) ResolveSailfish(verbose bool) {
	var (
		target  *Sailfish
		depth   uint
		visited []*Sailfish
	)

	depth = 0
	target = s
	visited = make([]*Sailfish, 0)

full_loop:
	for {
		if verbose {
			fmt.Println("* Checking for explosions")
		}

	explode_loop:
		for {
			if target.value == -1 {
				if depth == 4 {
					if target.left.leftWithValue != nil {
						target.left.leftWithValue.value += target.left.value
						target.left.leftWithValue.rightWithValue = target
					}

					if target.right.rightWithValue != nil {
						target.right.rightWithValue.value += target.right.value
						target.right.rightWithValue.leftWithValue = target
					}

					target.rightWithValue = target.right.rightWithValue
					target.leftWithValue = target.left.leftWithValue

					target.left = nil
					target.right = nil
					target.value = 0

					if target.rightWithValue != nil {
						target.rightWithValue.leftWithValue = target
					}
					if target.leftWithValue != nil {
						target.leftWithValue.rightWithValue = target
					}

					if verbose {
						fmt.Println("after explode: " + s.ToString())
					}

					target = s
					depth = 0
					visited = []*Sailfish{}
				}

				depth++
				target = target.left

				visited = append(visited, target)

				if verbose {
					fmt.Println("  Targetting "+target.ToString()+" at depth ", depth)
				}
			} else {
				if target.parent != nil {
					if target != target.parent.right {
						target = target.parent.right

						if verbose {
							fmt.Println("  Going right to "+target.ToString()+" at depth ", depth)
						}
					} else {
						target = target.parent.parent.right
						depth -= 1

						for {
							if verbose {
								fmt.Println("  Back up right "+target.ToString()+" at depth ", depth)
							}

							if target.parent == nil {
								break explode_loop
							}

							if !contains(visited, target) {
								break
							}

							if target.parent.parent == nil {
								break explode_loop
							} else {
								target = target.parent.parent.right
							}
							depth -= 1
						}
					}

					visited = append(visited, target)
				} else {
					break explode_loop
				}
			}
		}

		target = s
		depth = 0
		visited = []*Sailfish{}

		if verbose {
			fmt.Println("* Checking for splits")
		}

	split_loop:
		for {
			if target.value == -1 {
				depth++
				target = target.left

				visited = append(visited, target)

				if verbose {
					fmt.Println("  Targetting "+target.ToString()+" at depth ", depth)
				}
			} else if target.value >= 10 {
				target.left = NewSailfish(roundDown((float64(target.value)) / 2.0))
				target.left.parent = target
				target.left.leftWithValue = target.leftWithValue

				target.right = NewSailfish(roundUp((float64(target.value)) / 2.0))
				target.right.parent = target
				target.right.rightWithValue = target.rightWithValue

				target.left.rightWithValue = target.right
				target.right.leftWithValue = target.left
				if target.left.leftWithValue != nil {
					target.left.leftWithValue.rightWithValue = target.left
				}
				if target.right.rightWithValue != nil {
					target.right.rightWithValue.leftWithValue = target.right
				}

				target.value = -1
				target.rightWithValue = nil
				target.leftWithValue = nil

				if verbose {
					fmt.Println("after split " + s.ToString())
				}

				break split_loop
			} else {
				if target.parent != nil {
					if target != target.parent.right {
						target = target.parent.right

						if verbose {
							fmt.Println("  Going right to "+target.ToString()+" at depth ", depth)
						}
					} else {
						target = target.parent.parent.right
						depth -= 1

						for {
							if verbose {
								fmt.Println("  Back up right "+target.ToString()+" at depth ", depth)
							}

							if target.parent == nil {
								break full_loop
							}

							if !contains(visited, target) {
								break
							}

							if target.parent.parent == nil {
								break full_loop
							} else {
								target = target.parent.parent.right
							}
							depth -= 1
						}
					}

					visited = append(visited, target)
				} else {
					break full_loop
				}
			}
		}
	}
}

func (s *Sailfish) GetMagnitude() int {
	if s.value == -1 {
		return 3*s.left.GetMagnitude() + 2*s.right.GetMagnitude()
	}

	return s.value
}

func roundUp(val float64) int {
	return int(val + 0.6)
}

func roundDown(val float64) int {
	return int(val)
}

func contains(slice []*Sailfish, s *Sailfish) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}

	return false
}
