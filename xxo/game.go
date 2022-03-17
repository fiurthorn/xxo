package xxo

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame() *Game {
	return &Game{NewBoard()}
}

type Game struct {
	Board *Board
}

func (g *Game) rating(player *Player) int {
	if g.Board.Won() {
		factor := 1 + g.Board.Remaining()
		if player == g.Board.Winner() {
			return 10 * factor
		}
		return -10 * factor
	}
	return 0
}

func (g *Game) opposite(player *Player) *Player {
	if player == g.Board.player1 {
		return g.Board.player2
	}
	return g.Board.player1
}

//
// select the move with best case to win e.g. X..]
//                                       e.g. _O_]
//                                       e.g. ___]
//
func (g *Game) BestMove(player *Player) int {
	solutions := Solutions{[]int{}}
	_ = g.minimax(player, &solutions)

	index := rand.Intn(len(solutions.moves))
	return solutions.moves[index]
}

func (g *Game) minimax(player *Player, sol *Solutions) int {
	if g.Board.Won() || g.Board.Remaining() == 0 {
		return g.rating(player)
	}

	bestScore := -1000
	for i := 0; i < 9; i++ {
		if g.Board.IsEmpty(i) {
			g.Board.Set(i, player)
			score := -g.minimax(g.opposite(player), nil)
			g.Board.Reset(i)
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
