package server

import (
	"fmt"

	"github.com/tsrines/go-examples/deck"
)

// Run starts our api.
func Run() {
	fmt.Println("Got here")
	cards := deck.NewDeck()
	fmt.Println(cards)

}
