package day21

type DeterministicDice struct {
	lastValue int
	sides     int
}

func NewDeterministicDice() *DeterministicDice {
	dd := new(DeterministicDice)

	dd.lastValue = 0
	dd.sides = 100

	return dd
}

func (d *DeterministicDice) read() int {
	return d.lastValue
}

func (d *DeterministicDice) throw() int {
	d.lastValue = d.lastValue + 1
	if d.lastValue > d.sides {
		d.lastValue = 1
	}
	return d.lastValue
}
