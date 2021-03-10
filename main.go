package main

func main() {
	cards := deck{"Three of Hearts", newCard()}
	cards = append(cards, "Five of Spades")
	cards.print()

}

func deal() {

}

func newCard() string {
	return "Five of Diamonds"
}
