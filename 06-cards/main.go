package main

func main() {
	cards := newDeck()
	// cards.print()
	cards.saveToFile("cards.txt")
}
