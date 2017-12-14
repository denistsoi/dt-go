package main

func main() {
	cards := newDeck()
	// cards.print()
	hand, remaining := deal(cards, 2)

	hand.print()
	remaining.print()
}
