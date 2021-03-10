package main

import (
	"fmt"
)

func main() {
	nDeck := newDeck()
	nDeck.saveToFile("cards")
	newDeckFromFile("my")

	fmt.Println(len(nDeck))
}
