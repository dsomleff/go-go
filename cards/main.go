package main

func main() {
	cards := newDeck()

	hand, remainedDeck := deal(cards, 5)

	hand.print()
	remainedDeck.print()
}
