package main

import "fmt"

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
