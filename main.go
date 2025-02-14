package main

func main() {

	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Go will infer the type of the variable
	// card = "Five of Diamonds" // Reassign the value of the variable
	// card = newCard() // Call a function that returns a string

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// println("Hand of cards:")
	// hand.print()
	// println(" ")
	// println("Remaining cards:")
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting)) // Convert a string to a slice of bytes

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
