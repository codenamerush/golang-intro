package main

func main() {
	cards := newDeck()
	cards.saveToFile("myCards")
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	newDeckFromFile("myCards")
}
