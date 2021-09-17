package main

import (
	"fmt"
)

type Board struct {
	empty   *Player
	player1 *Player
	player2 *Player

	winner *Player
	fields [9]*Player
}

func NewBoard() *Board {
	empty := EmptyPlayer()
	return &Board{
		empty:   empty,
		player1: PlayerX(),
		player2: PlayerO(),
		fields: [9]*Player{
			empty, empty, empty,
			empty, empty, empty,
			empty, empty, empty,
		},
	}
}

func (b *Board) Winner() *Player {
	return b.winner
}

func (b *Board) Won() bool {
	return b.winner != nil && b.winner != b.empty
}

func (b *Board) Stopped() bool {
	return b.Remaining() == 0 || !b.Won()
}

func (b *Board) Remaining() int {
	count := 0
	for _, e := range b.fields {
		if e == b.empty {
			count++
		}
	}
	return count
}

// TODO
func (b *Board) Rating() {}

func (b *Board) coord(xy string) int {
	x := xy[0] - 'a'
	y := xy[1] - '1'
	return int(y*3 + x)
}

func (b *Board) index(x, y int) int {
	return (y-1)*3 + (x - 1)
}

func (b *Board) IsEmpty(x, y int) bool {
	return b.fields[b.index(x, y)] == b.empty
}

func (b *Board) IsEmptyS(xy string) bool {
	return b.fields[b.coord(xy)] == b.empty
}

func (b *Board) Get(x, y int) *Player {
	return b.fields[b.index(x, y)]
}

func (b *Board) GetS(xy string) *Player {
	return b.fields[b.coord(xy)]
}

func (b *Board) Set(x, y int, player *Player) {
	b.fields[b.index(x, y)] = player
}

func (b *Board) SetS(xy string, player *Player) {
	b.fields[b.coord(xy)] = player
}

func (b *Board) String() string {
	return fmt.Sprintf(` 3 | %s | %s | %s
---+---+---+---
 2 | %s | %s | %s
---+---+---+---
 1 | %s | %s | %s
---+---+---+---
   | a | b | c 
 `,
		b.fields[6], b.fields[7], b.fields[8],
		b.fields[5], b.fields[4], b.fields[3],
		b.fields[0], b.fields[1], b.fields[2],
	)
}
