// project Cards

// newDeck - Create a list of playing cards. Essentially an array of strings

// print - Log out the contents of the deck of cards

// shuffle - Shuffles cards in the deck

// deal - create a 'hand' of cards

// saveToFile - Save a list of cards to a file on the local machine

// newDeckFromFile - loads a list of cards from the local machine

package main

func main(){
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

