package main

// Player holds all the data about a single player
type Player struct {
	name string
	bank Bank
	deck []GemCard
}

// NewPlayer creates a player with an empty bank
func NewPlayer(name string) *Player {
	p := new(Player)
	p.name = name
	p.bank.Init()
	p.deck = []GemCard{}
	return p
}
