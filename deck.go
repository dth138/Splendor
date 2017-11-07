package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//GemCard is a playpiece in Splendor
type GemCard struct {
	cost  GemCost
	yield string
	vp    int
}

//GemCost is the purchase price of a card
type GemCost struct {
	black int
	white int
	red   int
	green int
	blue  int
}

//NewBoard creates a board state from a given Deck
func NewBoard(deck [][]GemCard) [][]GemCard {
	t1 := make([]GemCard, 0)
	t2 := make([]GemCard, 0)
	t3 := make([]GemCard, 0)
	tr := make([]GemCard, 0)
	for i := 0; i < 3; i++ {
		t1 = append(t1, deck[0][(len(deck[0])-1)])
		deck[0] = deck[0][:(len(deck[0]) - 1)]
		t2 = append(t2, deck[1][(len(deck[1])-1)])
		deck[1] = deck[1][:(len(deck[1]) - 1)]
		t3 = append(t3, deck[2][(len(deck[2])-1)])
		deck[2] = deck[2][:(len(deck[2]) - 1)]
		tr = append(tr, deck[3][(len(deck[3])-1)])
		deck[3] = deck[3][:(len(deck[3]) - 1)]
	}
	return [][]GemCard{t1, t2, t3, tr}
}

//PrettyPrintCard Prints a card in readable format
func PrettyPrintCard(c GemCard) {
	fmt.Println("_______________")
	fmt.Println("_" + c.yield + "      " + strconv.Itoa(c.vp))
	fmt.Println("_             _")
	fmt.Println("_ white     " + strconv.Itoa(c.cost.white) + " _")
	fmt.Println("_ blue      " + strconv.Itoa(c.cost.blue) + " _")
	fmt.Println("_ green     " + strconv.Itoa(c.cost.green) + " _")
	fmt.Println("_ red       " + strconv.Itoa(c.cost.red) + " _")
	fmt.Println("_ black     " + strconv.Itoa(c.cost.black) + " _")
	fmt.Println("_______________")
}

//PrettyPrintTier Prints a row of cards
func PrettyPrintTier(t []GemCard) {
	for _, c := range t {
		PrettyPrintCard(c)
	}
}

//ShuffleStack shuffles the remaining cards in each tier
func shuffleStack(d []GemCard) {
	for i := 0; i < 7; i++ { //Shuffling requires 7 shuffles
		for j := range d {
			shuf := rand.Intn(len(d))
			t := d[j]
			d[j] = d[shuf]
			d[shuf] = t
		}
	}
}

//ShuffleDeck shuffles all cards
func ShuffleDeck(d [][]GemCard) {
	for i := 0; i < 4; i++ {
		shuffleStack(d[i])
	}
}

// InitDeck Returns a new deck of Splendor Cards, of Type [][]GemCard
func InitDeck() [][]GemCard {
	t1 := make([]GemCard, 0)
	//Black, 8 Cards
	t1 = append(t1,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 4},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 2, red: 1, green: 0, blue: 2},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 1, red: 1, green: 1, blue: 2},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 3, blue: 0},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 0, red: 3, green: 1, blue: 0},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 2, red: 0, green: 2, blue: 0},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 1, green: 2, blue: 0},
			yield: "black"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 1, red: 1, green: 1, blue: 1},
			yield: "black"})

	//White, 8 Cards
	t1 = append(t1,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 4, blue: 0},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 3, red: 0, green: 0, blue: 1},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 0, red: 1, green: 2, blue: 1},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 2, white: 0, red: 0, green: 0, blue: 2},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 3},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 0, red: 1, green: 1, blue: 1},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 0, red: 2, green: 0, blue: 0},
			yield: "white"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 0, red: 0, green: 2, blue: 2},
			yield: "white"})

	//Red, 8 Cards
	t1 = append(t1,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 4, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 2, red: 2, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 1, red: 0, green: 1, blue: 1},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 2, red: 0, green: 1, blue: 1},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 2, white: 2, red: 0, green: 1, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 1, blue: 2},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 3, white: 1, red: 1, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 3, red: 0, green: 0, blue: 0},
			yield: "red"})

	//Green, 8 Cards
	t1 = append(t1,
		GemCard{vp: 1,
			cost:  GemCost{black: 4, white: 0, red: 0, green: 0, blue: 0},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 3, green: 0, blue: 0},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 2, white: 0, red: 2, green: 0, blue: 1},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 2, white: 1, red: 1, green: 0, blue: 1},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 1, red: 1, green: 0, blue: 1},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 1, red: 0, green: 1, blue: 3},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 2, green: 0, blue: 2},
			yield: "green"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 2, red: 0, green: 0, blue: 1},
			yield: "green"})

	//Blue, 8 Cards
	t1 = append(t1,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 0, red: 4, green: 0, blue: 0},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 2, white: 0, red: 0, green: 2, blue: 0},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 1, red: 2, green: 2, blue: 0},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 3, white: 0, red: 0, green: 0, blue: 0},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 1, green: 3, blue: 1},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 1, red: 1, green: 1, blue: 0},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 1, white: 1, red: 2, green: 1, blue: 0},
			yield: "blue"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 2, white: 1, red: 0, green: 0, blue: 0},
			yield: "blue"})

	t2 := make([]GemCard, 0)
	//Black 6 cards
	t2 = append(t2,
		GemCard{vp: 3,
			cost:  GemCost{black: 6, white: 0, red: 0, green: 0, blue: 0},
			yield: "black"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 0, red: 3, green: 5, blue: 0},
			yield: "black"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 5, red: 0, green: 0, blue: 0},
			yield: "black"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 0, red: 2, green: 4, blue: 1},
			yield: "black"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 2, white: 3, red: 0, green: 3, blue: 0},
			yield: "black"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 3, red: 0, green: 2, blue: 2},
			yield: "black"})

	//White 6 Cards
	t2 = append(t2,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 6, red: 0, green: 0, blue: 0},
			yield: "white"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 3, white: 0, red: 5, green: 0, blue: 0},
			yield: "white"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 0, red: 5, green: 0, blue: 0},
			yield: "white"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 2, white: 0, red: 4, green: 1, blue: 0},
			yield: "white"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 2, red: 3, green: 0, blue: 3},
			yield: "white"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 2, white: 0, red: 2, green: 3, blue: 0},
			yield: "white"})

	//Red 6 Cards
	t2 = append(t2,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 0, red: 6, green: 0, blue: 0},
			yield: "red"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 5, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 5, white: 3, red: 0, green: 0, blue: 0},
			yield: "red"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 1, red: 0, green: 2, blue: 4},
			yield: "red"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 3, white: 2, red: 2, green: 0, blue: 0},
			yield: "red"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 3, white: 0, red: 2, green: 0, blue: 3},
			yield: "red"})

	//Green 6 Cards
	t2 = append(t2,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 6, blue: 0},
			yield: "green"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 3, blue: 5},
			yield: "green"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 5, blue: 0},
			yield: "green"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 1, white: 4, red: 0, green: 0, blue: 2},
			yield: "green"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 2, white: 2, red: 0, green: 0, blue: 3},
			yield: "green"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 3, red: 3, green: 2, blue: 0},
			yield: "green"})

	//Blue 6 Cards
	t2 = append(t2,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 6},
			yield: "blue"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 4, white: 2, red: 1, green: 0, blue: 0},
			yield: "blue"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 5},
			yield: "blue"})
	t2 = append(t2,
		GemCard{vp: 2,
			cost:  GemCost{black: 0, white: 5, red: 0, green: 0, blue: 3},
			yield: "blue"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 3, white: 0, red: 0, green: 3, blue: 2},
			yield: "blue"})
	t2 = append(t2,
		GemCard{vp: 1,
			cost:  GemCost{black: 0, white: 0, red: 3, green: 2, blue: 2},
			yield: "blue"})

	t3 := make([]GemCard, 0)
	//Black 4 Cards
	t3 = append(t3,
		GemCard{vp: 5,
			cost:  GemCost{black: 3, white: 0, red: 7, green: 0, blue: 0},
			yield: "black"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 3, white: 0, red: 6, green: 3, blue: 0},
			yield: "black"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 0, white: 0, red: 7, green: 0, blue: 0},
			yield: "black"})
	t3 = append(t3,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 3, red: 3, green: 5, blue: 3},
			yield: "black"})

	//White 4 Cards
	t3 = append(t3,
		GemCard{vp: 5,
			cost:  GemCost{black: 7, white: 3, red: 0, green: 0, blue: 0},
			yield: "white"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 6, white: 3, red: 3, green: 0, blue: 0},
			yield: "white"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 7, white: 0, red: 0, green: 0, blue: 0},
			yield: "white"})
	t3 = append(t3,
		GemCard{vp: 3,
			cost:  GemCost{black: 3, white: 0, red: 5, green: 3, blue: 3},
			yield: "white"})

	//Red 4 Cards
	t3 = append(t3,
		GemCard{vp: 5,
			cost:  GemCost{black: 0, white: 0, red: 3, green: 7, blue: 0},
			yield: "red"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 0, white: 0, red: 3, green: 6, blue: 3},
			yield: "red"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 7, blue: 0},
			yield: "red"})
	t3 = append(t3,
		GemCard{vp: 3,
			cost:  GemCost{black: 3, white: 3, red: 0, green: 3, blue: 5},
			yield: "red"})

	//Green 4 Cards
	t3 = append(t3,
		GemCard{vp: 5,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 3, blue: 7},
			yield: "green"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 7},
			yield: "green"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 0, white: 3, red: 0, green: 3, blue: 6},
			yield: "green"})
	t3 = append(t3,
		GemCard{vp: 3,
			cost:  GemCost{black: 3, white: 5, red: 3, green: 0, blue: 3},
			yield: "green"})

	//Blue 4 Cards
	t3 = append(t3,
		GemCard{vp: 5,
			cost:  GemCost{black: 0, white: 7, red: 0, green: 3, blue: 0},
			yield: "blue"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 3, white: 6, red: 0, green: 0, blue: 3},
			yield: "blue"})
	t3 = append(t3,
		GemCard{vp: 4,
			cost:  GemCost{black: 0, white: 6, red: 0, green: 0, blue: 0},
			yield: "blue"})
	t3 = append(t3,
		GemCard{vp: 3,
			cost:  GemCost{black: 5, white: 3, red: 3, green: 3, blue: 0},
			yield: "blue"})

	//Royalty Tier, 10 cards
	tr := make([]GemCard, 0)
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 0, red: 4, green: 4, blue: 0},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 3, white: 3, red: 0, green: 0, blue: 3},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 4, blue: 4},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 3, red: 0, green: 3, blue: 3},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 4, red: 0, green: 0, blue: 4},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 3, white: 0, red: 3, green: 3, blue: 0},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 4, white: 0, red: 4, green: 0, blue: 0},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 4, white: 4, red: 0, green: 0, blue: 0},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 0, white: 0, red: 3, green: 3, blue: 3},
			yield: "none"})
	tr = append(tr,
		GemCard{vp: 3,
			cost:  GemCost{black: 3, white: 3, red: 3, green: 0, blue: 0},
			yield: "none"})

	return [][]GemCard{t1, t2, t3, tr}
}
