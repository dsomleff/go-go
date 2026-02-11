package main

import (
	"fmt"
	"math/rand"
	"os"
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

func (d deck) saveToFile(filename string) error  {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	s := strings.Split(string(bs), ",")
	return deck(s), nil
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
