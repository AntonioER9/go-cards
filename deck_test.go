package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Check the length of the deck
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Check the first card in the deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Check the last card in the deck
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove any files with the name "_decktesting"
	os.Remove("_decktesting")

	// Create a new deck
	deck := newDeck()

	// Save the deck to a file
	deck.saveToFile("_decktesting")

	// Load the deck from the file
	loadedDeck := newDeckFromFile("_decktesting")

	// Check the length of the loaded deck
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	// Remove the file "_decktesting"
	os.Remove("_decktesting")
}
