package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// first item slice must Spades of king
	if d[0] != "Spades of Ace" {
		t.Errorf("expected Spades of King but get %v", d[0])
	}

	// last item slice must Clubs of King
	if d[len(d) - 1] != "Clubs of King" {
		t.Errorf("Expected Clubs of King but get %v", d[len(d) - 1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// remove file _decktesting
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))	
	}

	os.Remove("_decktesting") 
}