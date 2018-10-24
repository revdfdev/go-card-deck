package main

// import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 7)

	hand.saveTofile("hand.txt")
	remainingDeck.saveTofile("remainindeck.txt")

}
