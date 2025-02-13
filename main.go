package main

func main() {

	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Go will infer the type of the variable
	// card = "Five of Diamonds" // Reassign the value of the variable
	// card = newCard() // Call a function that returns a string

	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
