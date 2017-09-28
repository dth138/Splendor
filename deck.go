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
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})
	t1 = append(t1,
		GemCard{vp: 0,
			cost:  GemCost{black: 0, white: 0, red: 0, green: 0, blue: 0},
			yield: "red"})

	t2 := make([]GemCard, 30)
	t3 := make([]GemCard, 20)
	return [][]GemCard{t1, t2, t3}
}
