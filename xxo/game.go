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
	return &Game{
		board: NewBoard(),
	}
}

type Game struct {
	board *Board

	fields [size]string
}

func (g *Game) ResetBoard() {
	g.board.ResetBoard()
	g.update()
}

func (g *Game) Winning() ([3]Pos, bool) {
	return g.board.Winning()
}

func (g *Game) Stopped() bool {
	return g.board.Stopped()
}

func (g *Game) IsEmpty(i int) bool {
	return g.board.IsEmpty(i)
}

func (g *Game) Get(i int) string {
	return g.fields[i]
}

func (g *Game) GetCurrent() *Player {
	return g.board.GetCurrent()
}

func (g *Game) Contains(pos [3]Pos, i int) bool {
	return g.board.Contains(pos, i)
}

func (g *Game) Set(i int, p *Player) {
	g.board.Set(i, p)
	g.update()
}

func (g *Game) update() {
	for i := 0; i < size; i++ {
		g.fields[i] = g.board.fields[i].Symbol()
	}
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
