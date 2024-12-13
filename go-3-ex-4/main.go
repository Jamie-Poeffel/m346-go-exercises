package main

import (
	"fmt"
)

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}


	// TODO: Loop over suits and ranks to output all combinations.
	for _, rank := range ranks {
		for _, suit := range suits {
			fmt.Printf("%c%c  ", suit, rank)
		}
		fmt.Println("")
	}
	// TODO: delete this line afterwards

	var cards map[rune][]string = map[rune][]string{}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards[suit] = append(cards[suit], fmt.Sprintf("%c", rank))
		}
	}

	fmt.Println("\nCards by suit:")
	for suit, cardList := range cards {
		fmt.Printf("%c: %v\n", suit, cardList)
	}}
