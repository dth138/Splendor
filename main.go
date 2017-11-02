package main

func main() {
	deck := InitDeck()

	ShuffleCards(deck[3])
	PrettyPrintTier(deck[3])
}
