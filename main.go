package main

func main() {
	deck := InitDeck()
	ShuffleDeck(deck)

	bank := NewHouseBank()
	playerOne := NewPlayer("Daniel")
	playerTwo := NewPlayer("Yimin")
	playerThree := NewPlayer("Karl")
	playerFour := NewPlayer("Julia")

	players := []Player{*playerOne, *playerTwo, *playerThree, *playerFour}

	RunSimulation(deck, *bank, players)
}
