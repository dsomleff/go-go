package main

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck len of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d) -1] != "Four of Clubs" {
		t.Errorf("Expected first card to be Four of Clubs, but got %v", d[len(d) -1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "_deck_testing")

	d := newDeck()

	err := d.saveToFile(filePath)
	if err != nil {
		t.Fatalf("Failed to save deck: %v", err)
	}


	loadedDeck, err := newDeckFromFile(filePath)
	if err != nil {
		t.Fatalf("Failed to load deck: %v", err)
	}

	if !reflect.DeepEqual(d, loadedDeck){
		t.Errorf("Loaded deck does not match saved deck")
	}
}
