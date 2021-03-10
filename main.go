package main

import "fmt"

func main() {
	newDeck := newDeck()
	hand, remandingDeck := deal(newDeck, 3)
	fmt.Println(hand, remandingDeck)

}

func newCard() string {
	return "Five of Diamonds"
}
