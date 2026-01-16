package main

func main() {
	cards := deck{"Ace of Spade", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "card name"
}
