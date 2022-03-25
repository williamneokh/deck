package Handler

import "testing"

func TestDealDeck(t *testing.T) {
	nDeck := NewDeck()
	var handSize = 1

	for handSize < 52 {

		d, _ := DealDeck(nDeck, handSize)

		if len(d) != handSize {

			t.Error("Wrong number of cards return", len(d))
		}

		handSize = handSize + 1
	}

}
