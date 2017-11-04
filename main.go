package main

func main() {
	deck := InitDeck()
	ShuffleDeck(deck)
	PrettyPrintTier(deck[3])
}
