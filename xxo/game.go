package xxo

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(),
		m:     &sync.Mutex{},
	}
}

type Game struct {
	board *Board

	fields [size]string
	m      sync.Locker
}

func (g *Game) Lock() {
	g.m.Lock()
}

func (g *Game) Unlock() {
	g.m.Unlock()
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

func (g *Game) Player1() *Player {
	return g.board.player1
}

func (g *Game) Player2() *Player {
	return g.board.player2
}

func (g *Game) IsEmpty(i int) bool {
	return g.board.IsEmpty(i)
}

func (g *Game) Get(i int) string {
	return g.fields[i]
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
		factor := 1 + g.board.Free()
		if player == g.board.Winner() {
			return 10 * factor
		}
		return -10 * factor
	}
	return 0
}

func (g *Game) Opposite(player *Player) *Player {
	if player == g.board.player1 {
		return g.board.player2
	}
	return g.board.player1
}

var possibleFields = map[int][][2]int{
	0: {{3, 6}, {4, 8}, {1, 2}},
	1: {{0, 2}, {4, 7}},
	2: {{0, 1}, {4, 6}, {5, 8}},
	3: {{0, 6}, {4, 5}},
	4: {{3, 5}, {0, 8}, {1, 7}, {2, 6}},
	5: {{3, 4}, {2, 8}},
	6: {{3, 0}, {4, 2}, {7, 8}},
	7: {{6, 8}, {4, 1}},
	8: {{6, 7}, {0, 4}, {2, 5}},
}

func (g *Game) adjust(solutions []int, p *Player) (selection []int) {
	opposite := g.Opposite(p)

	for _, idx := range solutions {
		lines := possibleFields[idx]
		for _, line := range lines {
			var count int
			a, b := g.board.Get(line[0]), g.board.Get(line[1])

			if a == g.board.empty || b == g.board.empty {
				count++
			}
			if a == opposite || b == opposite {
				count++
			}
			if count == 2 {
				selection = append(selection, idx)
			}
		}
	}

	return
}

func (g *Game) BestMove(player *Player) int {
	solutions := Solutions{[]int{}}
	_ = g.minimax(player, &solutions)

	log.Printf("move:%v", solutions.selection)
	selection := g.adjust(solutions.selection, player)
	length := len(selection)
	if length > 0 && length < len(solutions.selection) {
		log.Printf("shrink move:%v -> %v", solutions.selection, selection)
		solutions.selection = selection
	}

	index := rand.Intn(len(solutions.selection))
	return solutions.selection[index]
}

func (g *Game) minimax(player *Player, sol *Solutions) int {
	if g.board.Won() || g.board.Remaining() == 0 {
		return g.rating(player)
	}

	bestScore := -1000
	for i := 0; i < 9; i++ {
		if g.board.IsEmpty(i) {
			g.board.Set(i, player)
			score := -g.minimax(g.Opposite(player), nil)
			g.board.Reset(i)
			if score > bestScore {
				if sol != nil {
					sol.selection = []int{}
				}
				bestScore = score
			}
			if bestScore == score && sol != nil {
				sol.selection = append(sol.selection, i)
			}
		}
	}
	return bestScore
}

type Solutions struct {
	selection []int
}
