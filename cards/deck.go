package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

type deck []string

var separator string = ","

func newDeck() deck {
	var deck deck
	var cardSuits []string = []string{"S", "D", "H", "C"}
	var cardValues []string = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value+suite)
		}
	}

	return deck
}

func newDeckFromFile(filename string) (deck, error) {
	bs, error := ioutil.ReadFile(filename)
	return deck(strings.Split(string(bs), separator)), error
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) deal(handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join(d, separator)
}

func (d deck) print() {
	fmt.Println(len(d), "cards on this Deck[", d.toString(), "]")
}
