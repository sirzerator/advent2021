package day21

type Dice interface {
	throw() int
	read() int
}
