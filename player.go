package main

type Player struct {
	symbol string
}

func NewPlayer(symbol string) *Player {
	return &Player{
		symbol: symbol,
	}
}

func EmptyPlayer() *Player {
	return NewPlayer(" ")
}

func PlayerX() *Player {
	return NewPlayer("X")
}

func PlayerO() *Player {
	return NewPlayer("O")
}

func (p *Player) Symbol() string {
	return p.symbol
}

func (p *Player) Equals(o Player) bool {
	return p.symbol == o.symbol
}

func (p *Player) String() string {
	return p.symbol
}
