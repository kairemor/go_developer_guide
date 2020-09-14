package main

func main() {
	// var card = "Ace of Spades"
	// card := newCard() // defining new variable not assigning value to a varaiable
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_card")
	// fmt.Println(cards)

	// hand, remaimingDeck := deal(cards, 5)

	// hand.print()
	// remaimingDeck.print()
	// fmt.Println(cards.toString())

	//Conversion type a to type b
	// greeting := "Hi there"

	// fmt.Println([]byte(greeting))

}
