package main

type GemCard struct {
	cost GemCost
	vp   int
}

type GemCost struct {
	black int
	white int
	red   int
	green int
	blue  int
}

type Deck struct {
	deck [][]GemCard
}

func NewShuffledDeck() *Deck {
	return NewDeck()
}

func NewDeck() *Deck {
	d := new(Deck)
	return d
}

func InitDeck() [][]GemCard {
	t1 := make([]GemCard, 40)
	t2 := make([]GemCard, 30)
	t3 := make([]GemCard, 20)
	return [][]GemCard{t1, t2, t3}
}
