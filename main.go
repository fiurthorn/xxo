package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	board *Board
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0x77, 0x33, 0x11})
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
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

func (g *Game) minimax(player *Player, sol *Solutions) int {
	if g.board.Won() || g.board.Remaining() == 0 {
		return g.rating(player)
	}

	j := 0
	bestScore := -1000
	for i := 0; i < 8; i++ {
		// fmt.Printf("index: %d\n", i)
		//   var xy = board.pos(i);
		if g.board.IsEmpty(i) {
			j++
			g.board.Set(i, player)
			score := -g.minimax(g.opposite(player), nil)
			g.board.Reset(i)
			if score > bestScore {
				if sol != nil {
					for len(sol.moves) > 0 {
						sol.moves = []int{}
					}
				}
				bestScore = score
			}
			if bestScore == score && sol != nil {
				sol.moves = append(sol.moves, i)
				fmt.Printf("%v %v %v %v\n", i, bestScore, score, sol.moves)
			}
		}
	}
	fmt.Printf("j:%d\n", j)
	return bestScore
}

type Solutions struct {
	moves []int
}

func main() {
	g := &Game{
		board: NewBoard(),
	}
	// ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowTitle("Hello, World!")
	// if err := ebiten.RunGame(g); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(g.board)

	solutions := Solutions{[]int{}}
	fmt.Println(g.minimax(g.board.player1, &solutions))
	fmt.Println(solutions)
}
