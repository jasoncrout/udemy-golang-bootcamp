package main

import "fmt"

func main() {
	//cards := newDeck()

	// hand, remainingDeck := deal(cards, 4)

	// hand.print()
	// fmt.Println("-------------")
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()

	fmt.Println("-------")

	cards.shuffle()
	cards.print()
}
