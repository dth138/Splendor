package main

type GemCard struct {
	cost  GemCost
	yield string
	vp    int
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

	t2 := make([]GemCard, 30)
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

	t3 := make([]GemCard, 20)

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
	tr := make([]GemCard, 10)
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
