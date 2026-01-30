package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Diamonds", "Harts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Tree", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}
