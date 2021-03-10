package main

func main() {
	newDeck := newDeck()
	hand, remandingDeck := deal(newDeck, 3)
	hand.print()
	remandingDeck.print()

}
