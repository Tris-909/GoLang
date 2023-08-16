package main

import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// Open a new pack of cards
	newCards := newDeck()

	// Shuffle the new pack of cards
	cards := newCards.shuffle();

	// Use the shuffled pack of cards to deal
	dealedCards, newCards := deck.deal(cards, 5)

	fmt.Println(dealedCards)
	fmt.Println("Cards left:", newCards)

	saveCardsToFile(dealedCards);
	readCardsFromFile("test");
}
