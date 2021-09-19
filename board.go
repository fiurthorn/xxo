package xxo

import (
	"fmt"
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	size       = 3
	tileSize   = 80
	tileMargin = 4
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
	f         font.Face
)

func init() {
	var err error

	tileImage.Fill(color.White)
	goreg, err := opentype.Parse(goregular.TTF)
	if err != nil {
		fmt.Errorf("goregular: %w", err)
	}

	f, err = opentype.NewFace(goreg, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		fmt.Errorf("font face: %w", err)
	}
}

type Pos struct {
	X, Y int
}

type Board struct {
	empty   *Player
	player1 *Player
	player2 *Player

	fields [size * size]*Player
}

func NewBoard() *Board {
	empty := EmptyPlayer()
	return &Board{
		empty:   empty,
		player1: PlayerX(),
		player2: PlayerO(),
		fields: [size * size]*Player{
			empty, empty, empty,
			empty, empty, empty,
			empty, empty, empty,
		},
	}
}

func (b *Board) Size() (int, int) {
	x := size*tileSize + (size+1)*tileMargin
	y := x
	return x, y
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
		log.Println("game already ended!")
		return
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

func colorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	rf := float64(r) / 0xffff
	gf := float64(g) / 0xffff
	bf := float64(b) / 0xffff
	af := float64(a) / 0xffff
	// Convert to non-premultiplied alpha components.
	if 0 < af {
		rf /= af
		gf /= af
		bf /= af
	}
	return rf, gf, bf, af
}

func (bb *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			op := &ebiten.DrawImageOptions{}
			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			r, g, b, a := colorToScale(tileBackgroundColor)
			op.ColorM.Scale(r, g, b, a)
			boardImage.DrawImage(tileImage, op)

			player := bb.GetXY(i, j)
			if player != bb.empty {
				bound, _ := font.BoundString(f, player.Symbol())
				w := (bound.Max.X - bound.Min.X).Ceil()
				h := (bound.Max.Y - bound.Min.Y).Ceil()
				x = x + (tileSize-w)/2
				y = y + (tileSize-h)/2 + h
				text.Draw(boardImage, player.Symbol(), f, x, y, tileColor)
			}

		}
	}
}
