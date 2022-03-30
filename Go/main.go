package main

func main() {

	//slice - fixed array of items
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// typecasting
	// greeting := "Hi There!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my-cards")

	// cards := newDeckFromFile("my-cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
