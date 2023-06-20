package main

func main() {
	cards := newDeck()
	// cards.toString()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	

	// hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()

}
