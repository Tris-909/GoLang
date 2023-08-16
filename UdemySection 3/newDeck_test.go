package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newCards := newDeck()
	expectedNumberOfCards := 52
	if len(newCards) != expectedNumberOfCards {
		t.Errorf("Expected length of 52 but got %v", len(newCards))
	}

	expectedFirstCard := "1 of Spades";
	if newCards[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v but got %v", expectedFirstCard, newCards[0])
	}

	expectedLastCard := "K of Hearts";
	if newCards[expectedNumberOfCards-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v but got %v", expectedLastCard, newCards[expectedNumberOfCards-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("test");

	cards := newDeck();
	saveCardsToFile(cards);
	cardsFromFile := readCardsFromFile("test");
	expectedNumberOfCards := 52;

	if len(cardsFromFile) != expectedNumberOfCards {
		t.Errorf("Expecting %v cards from file but got %v", expectedNumberOfCards, len(cardsFromFile));
	}
}
