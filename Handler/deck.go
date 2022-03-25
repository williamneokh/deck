package Handler


import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func NewDeck() deck {

	cards := deck{}
	cardSuites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func (d deck) Print(message string) {
	fmt.Println("===" + message + "===")
	for i, card := range d {
		fmt.Println(i+1, card)
	}
	fmt.Println("===End of " + message + " ===")
}

func (d deck) ShuffleDeck() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d
}

func DealDeck(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func PrintDash() {
	fmt.Println("=======================")
}

func (d deck) SaveFile(filename string) error {

	bs := strings.Join(d, ",")

	err := os.WriteFile(filename, []byte(bs), 0666)

	if err != nil {
		log.Fatal("Unable to save file", err)
	}

	return err
}

func ReadFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	files := strings.Split(string(bs), ",")

	return files
}
