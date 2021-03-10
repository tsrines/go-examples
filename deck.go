package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Deck
type deck []string

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal is given a deck and a handSize.  We return slices of each (deck - handsize).
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// newDeck creates a 52 card deck, in theory.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardVals := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, vals := range cardVals {
			cards = append(cards, vals+" of "+suit)
		}
	}
	return cards
}

// toString converts the deck of cards into a CSV, ready for write.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
