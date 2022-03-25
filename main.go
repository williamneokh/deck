package main

import "log"

func main() {
	cardSet := NewDeck()
	cardSet.ShuffleDeck()

	player, dealer := DealDeck(cardSet, 10)

	err := player.SaveFile("playersCard")
	if err != nil {
		log.Fatal(err)
	}

	err = dealer.SaveFile("dealersCard")
	if err != nil {
		log.Fatal(err)
	}

	player = ReadFromFile("playersCard")
	player.Print("Player's Card")
	dealer = ReadFromFile("dealersCard")
	dealer.Print("Dealer's Card")

}
