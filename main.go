package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Let's win this game!\n")

	// hbank := NewHouseBank()
	d := NewShuffledDeck()
	fmt.Println(d)
}
