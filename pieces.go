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

type Bank struct {
	black int
	white int
	red   int
	green int
	blue  int
	gold  int
}

func NewHouseBank() *Bank {
	b := new(Bank)
	b.black = 7
	b.white = 7
	b.red = 7
	b.green = 7
	b.blue = 7
	b.gold = 5
	return b
}

func NewPlayerBank() *Bank {
	b := new(Bank)
	b.black = 0
	b.white = 0
	b.red = 0
	b.green = 0
	b.blue = 0
	b.gold = 0
	return b
}
