package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
	suit  string
	value string
}

func (c card) String() string {
	return c.value + " of " + c.suit
}

func cardFromString(c string) card {
	cSlice := strings.Split(c, " of ")

	return card{suit: cSlice[1], value: cSlice[0]}
}

type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, card{suit: cardSuit, value: cardValue})
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) StringSlice() []string {
	stringSlice := []string{}

	for _, card := range d {
		stringSlice = append(stringSlice, card.String())
	}

	return stringSlice
}

func (d deck) toString() string {
	return strings.Join(d.StringSlice(), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename+".txt", []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")
	d := []card{}

	for _, card := range stringSlice {
		d = append(d, cardFromString(card))
	}

	return d
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
