package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
)

type card struct {
	suit, rank string
}

type deck []card

func newDeck() deck {
	var cards deck
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, card{suit, rank})
		}
	}
	return cards
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range d {
		newPos := r.Intn(len(d))
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func (d *deck) drawCard() card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

func (d deck) print() {
	for _, card := range d {
		fmt.Printf("%s of %s\n", card.rank, card.suit)
	}
}

func (d deck) getValue() int {
	value := 0
	aceCount := 0

	for _, card := range d {
		switch card.rank {
		case "Ace":
			aceCount++
			value += 11
		case "Jack", "Queen", "King", "10":
			value += 10
		default:
			value += int(card.rank[0] - '0')
		}
	}

	for aceCount > 0 && value > 21 {
		value -= 10
		aceCount--
	}

	return value
}

func dealInitialCards(givenDeck *deck) (deck, deck) {
	var playerDeck, dealerDeck deck

	for i := 0; i < 2; i++ {
		playerDeck = append(playerDeck, givenDeck.drawCard())
		dealerDeck = append(dealerDeck, givenDeck.drawCard())
	}

	return playerDeck, dealerDeck
}

func hit(deck *deck, hand *deck) {
	*hand = append(*hand, deck.drawCard())
}

func playerTurn(deck *deck, playerHand *deck) {
	var choice string
	for choice != "s" {
		fmt.Println("\nYour hand:")
		playerHand.print()
		fmt.Printf("Total value: %d\n", playerHand.getValue())
		fmt.Print("Type 'h' to hit, type 's' to stand: \n")
		fmt.Scanln(&choice)

		if choice == "h" {
			hit(deck, playerHand)
			if playerHand.getValue() > 21 {
				fmt.Println("You busted!")
				os.Exit(0)
			}
		} else if choice != "s" {
			fmt.Println("Not a valid choice.")
		}
	}
}

func dealerTurn(deck *deck, dealerHand *deck) {
	for dealerHand.getValue() < 17 {
		fmt.Println()
		hit(deck, dealerHand)
	}
}

func compareHands(playerHand, dealerHand deck) {
	fmt.Println("\nYour hand:")
	playerHand.print()
	fmt.Printf("Total value: %d\n", playerHand.getValue())

	fmt.Println("\nDealer's hand:")
	dealerHand.print()
	fmt.Printf("Total value: %d\n", dealerHand.getValue())

	playerValue := playerHand.getValue()
	dealerValue := dealerHand.getValue()

	switch {
	case playerValue > 21:
		fmt.Println("Busted! You lose.")
	case dealerValue > 21:
		fmt.Println("Dealer busted! You win.")
	case playerValue == dealerValue:
		fmt.Println("It's a tie!")
	case playerValue > dealerValue:
		fmt.Println("You win!")
	default:
		fmt.Println("You lose.")
	}
}

func main() {
	fmt.Println("Welcome to Blackjack!")
	fmt.Println("---------------------")

	deck := newDeck()
	deck.shuffle()

	playerHand, dealerHand := dealInitialCards(&deck)

	playerTurn(&deck, &playerHand)
	dealerTurn(&deck, &dealerHand)

	compareHands(playerHand, dealerHand)
}
