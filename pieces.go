package main

//Bank holds the resources of the bank
type Bank struct {
	black int
	white int
	red   int
	green int
	blue  int
	gold  int
}

//NewHouseBank initializes the game pieces to begin play
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

//NewPlayerBank creates an empty player inventory
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
