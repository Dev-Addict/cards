package main

func main() {
	cards := newDeck()

	cards.saveToFile("cards")

	newDeckFromFile("cards.txt").print()
}
