package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // _ is a placeholder for a variable that we don't intend to use
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() { // (d deck) is a receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // Return two values of type deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// []string(d) // Convert deck to a slice of strings
	return strings.Join([]string(d), ",") // Convert slice of strings to a single string
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil { // If there is an error
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// string(bs) // Convert byte slice to string
	s := strings.Split(string(bs), ",") // Convert string to slice of strings
	return deck(s)                      // Convert slice of strings to deck
}
