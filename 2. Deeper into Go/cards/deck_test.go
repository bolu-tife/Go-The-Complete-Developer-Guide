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

	if d[0] != "A of Spades" {
		t.Errorf("Expected first card of A of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected last card of K of Clubs, but got %v", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fileName := "temp_deck_testing"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 52{
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))

	}
	os.Remove(fileName)
}
