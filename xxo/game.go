package xxo

import (
	"image/color"
	"log"
	"math/rand"
	"time"
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

type Game struct {
	board *Board
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
