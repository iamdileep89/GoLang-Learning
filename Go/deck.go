package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck which is slice of strings
type deck []string

// function to create new deck of cards
func newDeck() deck {
	cards := deck{}                                                // empty slice of type deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //slice of strings
	cardValues := []string{"Ace", "Two", "Three", "Four"}          // slice of strings

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function for deck to print items
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// a regular function to get a portion of a deck or slice of strings
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function for deck to join all decks seperated by ","
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//receiver function for deck to save a string to a file in the form of ASCII bytes
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//regular function to read data from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //bs - byteSlice
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//receiver function to shuffle the cards with the help of random number generator inbuilt functions
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
