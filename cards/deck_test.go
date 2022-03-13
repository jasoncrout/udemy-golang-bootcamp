package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	//test deck has 52 cards
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	//test first card is ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	//test last card is king of clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 Cards in deck but found %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
