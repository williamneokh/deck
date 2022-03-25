package main

import (
	"log"

	"github.com/williamneokh/Deck/Handler"
)

func main() {
	cardSet := Handler.NewDeck()
	cardSet.ShuffleDeck()

	player, dealer := Handler.DealDeck(cardSet, 10)

	err := player.SaveFile("playersCard")
	if err != nil {
		log.Fatal(err)
	}

	err = dealer.SaveFile("dealersCard")
	if err != nil {
		log.Fatal(err)
	}

	player = Handler.ReadFromFile("playersCard")
	player.Print("Player's Card")
	dealer = Handler.ReadFromFile("dealersCard")
	dealer.Print("Dealer's Card")

}
