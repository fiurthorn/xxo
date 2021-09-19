package xxo

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
)

var (
	backgroundColor     = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor          = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	tileBackgroundColor = color.RGBA{0xee, 0xe4, 0xda, 0xff}
	tileColor           = color.RGBA{0x77, 0x6e, 0x65, 0xff}
)

func NewGame(input *Input, board *Board) *Game {
	g := Game{
		input: input,
		board: board,
	}
	g.player = g.board.player1
	return &g
}

type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image

	player *Player
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.Input(g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Input(i *Input) error {
	if i.reset {
		i.reset = false
		for i := range g.board.fields {
			g.board.fields[i] = g.board.empty
		}
	}
	if i.isReleased {
		i.isReleased = false

		sw, sh := ebiten.WindowSize()
		if sw == 0 || sh == 0 {
			sw = ScreenWidth
			sh = ScreenHeight
		}
		bw, bh := g.board.Size()
		x := (sw - bw) / 2
		y := (sh - bh) / 2

		// log.Printf("sw(%d), sh(%d)", sw, sh)
		// log.Printf("bw(%d), bh(%d)", bw, bh)

		// log.Printf("x(%d) := (sw - bw) / 2", x)
		// log.Printf("y(%d) := (sh - bh) / 2", y)

		i, j := (-1*(x-i.x))/tileSize, (-1*(y-i.y))/tileSize

		// log.Printf("i(%d) %d", i, tileSize)
		// log.Printf("j(%d) %d", j, tileSize)

		if i < 0 || j < 0 || i >= 3 || j >= 3 {
			return nil
		}
		if (g.board.Won() || g.board.Remaining() == 0) && i == j && i == 1 {
			g.input.reset = true
			return nil
		}
		if g.board.IsEmptyXY(i, j) {
			g.board.SetXY(i, j, g.player)

			if g.board.Remaining() > 0 && !g.board.Won() {
				ai := g.opposite(g.player)
				index := g.BestMove(ai)
				g.board.Set(index, ai)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage = ebiten.NewImage(w, h)
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2

	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("outsideWidth:%d, outsideHeight:%d", outsideWidth, outsideHeight)
	// log.Printf("ScreenWidth:%d, ScreenHeight:%d", ScreenWidth, ScreenHeight)
	return ScreenWidth, ScreenHeight
}

func (g *Game) rating(player *Player) int {
	if g.board.Won() {
		factor := 1 + g.board.Remaining()
		if player == g.board.Winner() {
			return 10 * factor
		}
		return -10 * factor
	}
	return 0
}

func (g *Game) opposite(player *Player) *Player {
	if player == g.board.player1 {
		return g.board.player2
	}
	return g.board.player1
}

func (g *Game) BestMove(player *Player) int {
	solutions := Solutions{[]int{}}
	_ = g.minimax(player, &solutions)

	index := rand.Intn(len(solutions.moves))
	return solutions.moves[index]
}

func (g *Game) minimax(player *Player, sol *Solutions) int {
	if g.board.Won() || g.board.Remaining() == 0 {
		return g.rating(player)
	}

	bestScore := -1000
	for i := 0; i < 9; i++ {
		if g.board.IsEmpty(i) {
			g.board.Set(i, player)
			score := -g.minimax(g.opposite(player), nil)
			g.board.Reset(i)
			if score > bestScore {
				if sol != nil {
					sol.moves = []int{}
				}
				bestScore = score
			}
			if bestScore == score && sol != nil {
				sol.moves = append(sol.moves, i)
			}
		}
	}
	return bestScore
}

type Solutions struct {
	moves []int
}
