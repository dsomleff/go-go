package main

import "fmt"

func main() {
	card := newCard("Five of Diamonds")

	fmt.Printf(card)
}

func newCard(cardName string) string {
	return cardName
}

