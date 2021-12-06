package day4

import (
	"fmt"
	"strings"

	d "github.com/sirzerator/advent2021/days"
)

type Bingo struct {
	board      [][]int
	lastCalled int
	winner     bool
}

func NewBingo(board []string) *Bingo {
	b := new(Bingo)
	b.winner = false
	b.lastCalled = 0
	b.board = [][]int{
		d.ArrayToInteger(strings.Fields(board[0])),
		d.ArrayToInteger(strings.Fields(board[1])),
		d.ArrayToInteger(strings.Fields(board[2])),
		d.ArrayToInteger(strings.Fields(board[3])),
		d.ArrayToInteger(strings.Fields(board[4])),
	}
	return b
}

func (b *Bingo) EvaluateScore() int {
	sum := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			sum += b.board[i][j]
		}
	}

	return sum * b.lastCalled
}

func (b *Bingo) MarkAndTestIsWinner(x int) bool {
	b.Mark(x)
	b.winner = b.IsWinner()
	return b.winner
}

func (b *Bingo) Mark(x int) {
	b.lastCalled = x

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.board[i][j] == x {
				b.board[i][j] = 0
				return
			}
		}
	}
}

func (b *Bingo) IsWinner() bool {
	// Rows
	for i := 0; i < 5; i++ {
		if b.board[i][0]+b.board[i][1]+b.board[i][2]+b.board[i][3]+b.board[i][4] == 0 {
			return true
		}
	}

	// Columns
	for j := 0; j < 5; j++ {
		if b.board[0][j]+b.board[1][j]+b.board[2][j]+b.board[3][j]+b.board[4][j] == 0 {
			return true
		}
	}

	return false
}

func (b *Bingo) ToString() string {
	return fmt.Sprintf(
		"%3d %3d %3d %3d %3d\n%3d %3d %3d %3d %3d\n%3d %3d %3d %3d %3d\n%3d %3d %3d %3d %3d\n%3d %3d %3d %3d %3d",
		b.board[0][0],
		b.board[0][1],
		b.board[0][2],
		b.board[0][3],
		b.board[0][4],
		b.board[1][0],
		b.board[1][1],
		b.board[1][2],
		b.board[1][3],
		b.board[1][4],
		b.board[2][0],
		b.board[2][1],
		b.board[2][2],
		b.board[2][3],
		b.board[2][4],
		b.board[3][0],
		b.board[3][1],
		b.board[3][2],
		b.board[3][3],
		b.board[3][4],
		b.board[4][0],
		b.board[4][1],
		b.board[4][2],
		b.board[4][3],
		b.board[4][4],
	)
}
