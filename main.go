package main

import (
	"fmt"
)

func main() {
	nDeck := newDeck()
	nDeck.saveToFile("cards")

	fmt.Println(len(nDeck))
}
