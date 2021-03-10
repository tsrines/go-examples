package main

import "fmt"

// Deck
type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardVals := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, vals := range cardVals {
			cards = append(cards, vals+" of "+suit)
		}
	}

	fmt.Println(cards)
	return cards
}
