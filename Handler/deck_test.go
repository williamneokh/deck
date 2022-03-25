package Handler

import "testing"

func TestDealDeck(t *testing.T) {
	nDeck := NewDeck()
	var handSize = 6

	for handSize < 1 {

		d, _ := DealDeck(nDeck, handSize)

		if len(d) != handSize {

			t.Error("Wrong number of cards return", d)
		}

		handSize = handSize + 1
	}

}
