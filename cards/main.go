package main

func main() {
  cards := newDeckFromFile("my_cards")
  cards.print()

  // testing dealing a hand
  hand, remainingDeck := deal(cards, 5)
  hand.print()
  remainingDeck.print()

  // testing converting a deck to a string
  // fmt.Println(cards.toString())

  // saving deck to file
  // cards.saveToFile("my_cards")
}