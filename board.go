package main

import (
	"fmt"
)

type Board struct {
	empty   *Player
	player1 *Player
	player2 *Player

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

func (b *Board) inALine(k, l, m Pos) bool {
	K, L, M := b.GetPos(k), b.GetPos(l), b.GetPos(m)
	return K == L && L == M && M != b.empty
}

func (b *Board) Winner() *Player {
	if b.inALine(Pos{0, 0}, Pos{1, 1}, Pos{2, 2}) {
		return b.GetXY(1, 1)
	}

	if b.inALine(Pos{2, 0}, Pos{1, 1}, Pos{0, 2}) {
		return b.GetXY(1, 1)
	}

	for i := 0; i < 3; i++ {
		if b.inALine(Pos{0, i}, Pos{1, i}, Pos{2, i}) {
			return b.GetXY(i, i)
		}
		if b.inALine(Pos{i, 0}, Pos{i, 1}, Pos{i, 2}) {
			return b.GetXY(i, i)
		}
	}

	return b.empty
}

func (b *Board) Won() bool {
	winner := b.Winner()
	return winner != nil && winner != b.empty
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

func (b *Board) byCoord(xy string) int {
	x := int(xy[0] - 'a')
	y := int(xy[1] - '1')
	if x < 0 || y < 0 || x > 2 || y > 2 {
		panic(fmt.Sprintf("pos out of range: %#v", xy))
	}
	return b.byXY(x, y)
}

func (b *Board) byPos(p Pos) int {
	if p.X < 0 || p.Y < 0 || p.X > 2 || p.Y > 2 {
		panic(fmt.Sprintf("pos out of range: %#v", p))
	}
	return b.byXY(p.Y, p.X)
}

func (b *Board) byIndex(index int) int {
	if index < 0 || index > 8 {
		panic(fmt.Sprintf("index out of range %+v", index))
	}
	return index
}

func (b *Board) byXY(x, y int) int {
	if x < 0 || y < 0 || x > 2 || y > 2 {
		panic(fmt.Sprintf("pos out of range: %#v, %#v", x, y))
	}
	return y*3 + x
}

func (b *Board) IsEmpty(index int) bool {
	return b.fields[b.byIndex(index)] == b.empty
}

func (b *Board) IsEmptyPos(p Pos) bool {
	return b.IsEmpty(b.byPos(p))
}

func (b *Board) IsEmptyXY(x, y int) bool {
	return b.IsEmpty(b.byXY(x, y))
}

func (b *Board) IsEmptyCoord(xy string) bool {
	return b.IsEmpty(b.byCoord(xy))
}

func (b *Board) Get(index int) *Player {
	return b.fields[b.byIndex(index)]
}

func (b *Board) GetXY(x, y int) *Player {
	return b.Get(b.byXY(x, y))
}

func (b *Board) GetPos(p Pos) *Player {
	return b.Get(b.byPos(p))
}

func (b *Board) GetCoord(xy string) *Player {
	return b.Get(b.byCoord(xy))
}

func (b *Board) Reset(index int) {
	b.fields[b.byIndex(index)] = b.empty
}

func (b *Board) Set(index int, player *Player) {
	if b.Won() {
		panic("game already ended!")
	}
	b.fields[b.byIndex(index)] = player
}

func (b *Board) SetPos(p Pos, player *Player) {
	b.Set(b.byPos(p), player)
}

func (b *Board) SetXY(x, y int, player *Player) {
	b.Set(b.byXY(x, y), player)
}

func (b *Board) SetCoord(xy string, player *Player) {
	b.Set(b.byCoord(xy), player)
}

func (b *Board) String() string {
	return fmt.Sprintf(`
 3 | %s | %s | %s
---+---+---+---
 2 | %s | %s | %s
---+---+---+---
 1 | %s | %s | %s
---+---+---+---
   | a | b | c 
 `,
		b.fields[6], b.fields[7], b.fields[8],
		b.fields[3], b.fields[4], b.fields[5],
		b.fields[0], b.fields[1], b.fields[2],
	)
}
