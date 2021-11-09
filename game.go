package xxo

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ScreenWidth  = 420
	ScreenHeight = 600
)

var (
	backgroundColor     = color.RGBA{R: 0xfa, G: 0xf8, B: 0xef, A: 0xff}
	frameColor          = color.RGBA{R: 0xbb, G: 0xad, B: 0xa0, A: 0xff}
	activeColor         = color.RGBA{G: 0x77, B: 0xaa, A: 0xff}
	tileBackgroundColor = color.RGBA{R: 0xee, G: 0xe4, B: 0xda, A: 0xff}
	tileColor           = color.RGBA{R: 0x77, G: 0x6e, B: 0x65, A: 0xff}
)

func NewGame(input *Input, board *Board) *Game {
	g := Game{
		input: input,
		board: board,
	}
	g.currentPlayer = g.board.player1
	return &g
}

type Game struct {
	input *Input
	board *Board

	newGameButton *Button
	currentPlayer *Player
}

func (g *Game) Update() error {
	g.input.Update()
	if err := g.Input(g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Input(inp *Input) error {
	if inp.reset {
		inp.reset = false
		for i := range g.board.fields {
			g.board.fields[i] = g.board.empty
		}
	}
	if inp.isReleased {
		inp.isReleased = false

		sw, sh := ebiten.WindowSize()
		if sw == 0 || sh == 0 {
			sw = ScreenWidth
			sh = ScreenHeight
		}
		bw, bh := g.board.Size()
		x := (sw - bw) / 2
		y := (sh - bh) / 2

		{
			bw, bh := g.board.boardImage.Size()
			x1, y1 := (sw-bw)/2, (sh-bh)/2+bh+10
			x2, y2 := x1+bw, y1+tileSize/2

			if g.board.Won() || g.board.Remaining() == 0 && inp.x > x1 && inp.x < x2 && inp.y > y1 && inp.y < y2 {
				g.input.reset = true
				return nil
			}
		}

		i, j := -1*(x-inp.x), -1*(y-inp.y)
		if i < 0 || j < 0 || i > bw || j > bh {
			return nil
		}
		i, j = i/tileSize, j/tileSize

		if i < 0 || j < 0 || i >= 3 || j >= 3 {
			return nil
		} else if g.board.IsEmptyXY(i, j) {
			g.board.SetXY(i, j, g.currentPlayer)

			if g.board.Remaining() > 0 && !g.board.Won() {
				ai := g.opposite(g.currentPlayer)
				index := g.BestMove(ai)
				g.board.Set(index, ai)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	g.board.Draw()
	g.DrawBoard(screen)
	g.DrawNewGameButton(screen)
}

func (g *Game) DrawBoard(screen *ebiten.Image) {
	sw, sh := screen.Size()
	bw, bh := g.board.boardImage.Size()
	x, y := (sw-bw)/2, (sh-bh)/2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.board.boardImage, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if outsideWidth != ScreenWidth || outsideHeight != ScreenHeight {
		ScreenWidth = outsideWidth
		ScreenHeight = outsideHeight
	}
	return outsideWidth, outsideHeight
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
			if sol != nil {
				log.Printf("index %d: score %d", i, score)
			}
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
