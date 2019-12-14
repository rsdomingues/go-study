package main

import "fmt"

func main() {
	filename := "deck.txt"

	// Generate a new deck file from the disk. If something goes wrong create a brand new one
	deck, error := newDeckFromFile(filename)
	if nil != error {
		fmt.Println("Cloud not read file")
		deck = newDeck()
	}
	deck.print()

	//Shufling the deck
	deck.shuffle()
	deck.print()

	//Deal a simple hand
	firstHand, remainingDeck := deck.deal(5)
	firstHand.print()
	remainingDeck.print()

	//save the newly created deck to file
	deck.saveToFile(filename)
}
