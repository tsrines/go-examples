package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck hello goodbye.
type Deck []string

// NewDeckFromFile retrieves the byte slice of the filename, and returns a Deck.
func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
		newDeck()
	}
	return Deck(strings.Split(string(bs), ","))
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}

// deal is given a Deck and a handSize.  We return slices of each (Deck - handsize).
func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// NewDeck creates a 52 card Deck, in theory.
func NewDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardVals := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, vals := range cardVals {
			cards = append(cards, vals+" of "+suit)
		}
	}
	return cards
}

// ToString converts the Deck of cards into a CSV, ready for write.
func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}
