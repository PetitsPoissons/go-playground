package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Create a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♥️", "♠️", "♦️" ,"♣️"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// Print cards from a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal two hands of variable hand size from a deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Take a deck type and turn it into a string
func (d deck) toString() string {
	// convert d to a slice of strings with []string(d)
	// turn slice of strings into a single string
	return strings.Join([]string(d), ",")
}

// Save to a deck to a file (permission 0666 means anyone can read and write the file)
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Retrieve deck from file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// convert the byte slice bs into a string, then split it by comma and convert into slice of strings
	ss := strings.Split(string(bs), ",")
	// convert the slice of strings into a deck and return it
	return deck(ss)
}

// Shuffle a deck of cards
func (d deck) shuffle() {
	// randomize the seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// shuffle the deck
	for i := range d {
		// new position is a random nb between 0 and (length of slice - 1)
		newPosition := r.Intn(len(d) - 1)
		// swap positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}