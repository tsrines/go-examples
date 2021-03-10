package main

import "fmt"

func main() {
	newDeck := newDeck()
	fmt.Println(newDeck)

}

func newCard() string {
	return "Five of Diamonds"
}
