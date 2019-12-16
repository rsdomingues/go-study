package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck to have 52 cards, but got: %v", len(d))
	}

	if d[0] != "AS" {
		t.Errorf("Expected the first card of the deck to be AS, but got: %v", d[0])
	}
	if d[len(d)-1] != "KC" {
		t.Errorf("Expected the last card of the deck to be AS, but got: %v", d[len(d)-1])
	}
}

func TestToString(t *testing.T) {
	ds := newDeck().toString()

	if ds != "AS,2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS,AD,2D,3D,4D,5D,6D,7D,8D,9D,10D,JD,QD,KD,AH,2H,3H,4H,5H,6H,7H,8H,9H,10H,JH,QH,KH,AC,2C,3C,4C,5C,6C,7C,8C,9C,10C,JC,QC,KC" {
		t.Errorf("Expected to string format not used, instead found this: %v", ds)
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	h, nd := d.deal(7)

	if len(h) != 7 {
		t.Errorf("Expected hand to have 7 cards, but got: %v", len(h))
	}

	if len(nd) != len(d)-7 {
		t.Errorf("Expected new deck to have 45 cards, but got: %v", len(nd))
	}

	if len(d) != 52 {
		t.Errorf("Expected deck to have 52 cards, but got: %v", len(d))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	d.shuffle()

	ds := d.toString()

	if ds == "AS,2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS,AD,2D,3D,4D,5D,6D,7D,8D,9D,10D,JD,QD,KD,AH,2H,3H,4H,5H,6H,7H,8H,9H,10H,JH,QH,KH,AC,2C,3C,4C,5C,6C,7C,8C,9C,10C,JC,QC,KC" {
		t.Errorf("Expected suffle, to change the default structure, istead got: %v", ds)
	}
}

func TestSaveToFileAndNewFromFile(t *testing.T) {
	d := newDeck()
	d.shuffle()
	filename := "_testingsavedeck.file"

	os.Remove(filename)

	error := d.saveToFile(filename)
	fileinfo, err := os.Stat(filename)

	if nil != error {
		t.Errorf("Expected to save file, but got error instead, %v", error.Error())
	}

	if nil != err && fileinfo.Name() != filename {
		t.Errorf("Expected file %v to exist, but file not found.", filename)
	}

	nd, error := newDeckFromFile(filename)

	if nil != error {
		t.Errorf("To load an entire deck from file, instead got error %v", error.Error())
	}

	if len(nd) != len(d) {
		t.Errorf("expected to load the same deck from file, but got diferent sizes %v", len(nd))
	}

	for i, card := range d {
		if nd[i] != card {
			t.Errorf("expected to load the same deck from file, but got diferent card on position %v, expected %v but got %v", i, card, nd[i])
		}
	}

	os.Remove(filename)

}
