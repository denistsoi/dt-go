package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("cards.txt")
	cards := newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
}
