package xxo

import (
	"fmt"
	"log"
)

const (
	side = 3
	size = side * side
)

type Pos struct {
	X, Y int
}

type Board struct {
	empty   *Player
	player1 *Player
	player2 *Player
	current *Player

	fields [size]*Player
}

func NewBoard() *Board {
	empty := EmptyPlayer()
	b := Board{
		empty:   empty,
		player1: PlayerX(),
		player2: PlayerO(),
		fields: [size]*Player{
			empty, empty, empty,
			empty, empty, empty,
			empty, empty, empty,
		},
	}
	b.current = b.player1
	return &b
}

func (b *Board) inALine(p [side]Pos) bool {
	K, L, M := b.GetPos(p[0]), b.GetPos(p[1]), b.GetPos(p[2])
	return K == L && L == M && M != b.empty
}

var lines = [8][side]Pos{
	{{0, 0}, {1, 1}, {2, 2}},
	{{2, 0}, {1, 1}, {0, 2}},

	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},

	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
}

func (b *Board) Winning() ([side]Pos, bool) {
	for _, p := range lines {
		if b.inALine(p) {
			return p, true
		}
	}
	return [side]Pos{}, false
}

func (b *Board) Winner() *Player {
	if line, ok := b.Winning(); ok {
		return b.GetPos(line[0])
	}
	return b.empty
}

func (b *Board) Won() bool {
	winner := b.Winner()
	return winner != nil && winner != b.empty
}

func (b *Board) Stopped() bool {
	return b.Remaining() == 0 || b.Won()
}

func (b *Board) Remaining() (count int) {
	if b.Won() {
		return
	}

	for _, e := range b.fields {
		if e == b.empty {
			count++
		}
	}
	return
}

func (b *Board) Contains(line [side]Pos, i int) bool {
	for _, idx := range line {
		if i == b.byPos(idx) {
			return true
		}
	}

	return false
}

func (b *Board) byPos(p Pos) int {
	if p.X < 0 || p.Y < 0 || p.X >= side || p.Y >= side {
		panic(fmt.Sprintf("pos out of range: %#v", p))
	}
	return p.Y*side + p.X
}

func (b *Board) byIndex(index int) int {
	if index < 0 || index >= size {
		panic(fmt.Sprintf("index out of range %+v", index))
	}
	return index
}

func (b *Board) IsEmpty(index int) bool {
	return b.fields[b.byIndex(index)] == b.empty
}

func (b *Board) IsEmptyPos(p Pos) bool {
	return b.IsEmpty(b.byPos(p))
}

func (b *Board) Get(index int) *Player {
	return b.fields[b.byIndex(index)]
}

func (b *Board) GetPos(p Pos) *Player {
	return b.Get(b.byPos(p))
}

func (b *Board) GetPlayerX() string {
	return b.player1.symbol
}

func (b *Board) GetPlayerO() string {
	return b.player2.symbol
}

func (b *Board) GetCurrent() *Player {
	return b.current
}

func (b *Board) toggle() {
	if b.current == b.player1 {
		b.current = b.player2
		return
	}
	b.current = b.player1
}

func (b *Board) Reset(index int) {
	b.fields[b.byIndex(index)] = b.empty
}

func (b *Board) ResetBoard() {
	for i := 0; i < size; i++ {
		b.fields[b.byIndex(i)] = b.empty
	}
}

func (b *Board) Set(index int, player *Player) {
	if b.Won() {
		log.Println("game already ended!")
		return
	}
	b.fields[b.byIndex(index)] = player
	b.toggle()
}

func (b *Board) SetPos(p Pos, player *Player) {
	b.Set(b.byPos(p), player)
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
