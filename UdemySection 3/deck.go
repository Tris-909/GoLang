package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"of Spades", "of Clovers", "of Diamonds", "of Hearts"};
	cardsValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, cardSuit := range cardsSuits {
		for _, cardValue := range cardsValues {
			cards = append(cards, cardValue + " " + cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(numberOfCardNeedToBeDeal int) (deck, deck) {
	numberOfCardLefts := len(d);
	newDealedCards := deck{}
	for i := 0; i < numberOfCardNeedToBeDeal; i++ {
		// Call random for a number in cards
		currentCardNumber := rand.Intn(numberOfCardLefts+1);
		// Add that random card to the new deal hand
		newDealedCards = append(newDealedCards, d[currentCardNumber]);
		// Remove the card you have dealed from the deck
		d = removeOneCardFromCards(d, currentCardNumber);
		numberOfCardLefts = numberOfCardLefts -1;
	}

	fmt.Println("Cards left in the deck", d);
	// Return d also so we have the remaining cards for another user to deal
	return newDealedCards, d;
}

func removeOneCardFromCards(cards deck, index int) deck {
	newDecks := append(cards[:index], cards[index+1:]...);

	return newDecks;
}

func saveCardsToFile(cards deck) {
	cardString := strings.Join(cards, ",");
	bytesOfCards := []byte(cardString);
	ioutil.WriteFile("test", bytesOfCards, 0644);
}

func readCardsFromFile(filePath string) deck {
	content, err := ioutil.ReadFile(filePath);
	if (err != nil) {
		log.Fatal(err);
		os.Exit(1);
	}

	cardsFromFile := strings.Split(string(content), ",");
	return cardsFromFile;
}

func (d deck) shuffle() deck {
	for i, _ := range d {
		// Call random for a number in cards
		currentSwitchingPosition := rand.Intn(len(d)+1-i);

		currentIndexReference := d[i];
		d[i] = d[currentSwitchingPosition];
		d[currentSwitchingPosition] = currentIndexReference;

	}
	return d;
}