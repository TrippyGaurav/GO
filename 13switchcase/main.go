package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Introduction to switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice : 1 and you can open")
		fallthrough

	case 2:
		fmt.Println("You can move to  2 spot")

	case 3:
		fmt.Println("You can move to 3 spot")

	case 4:
		fmt.Println("You can move to 4 spot")

	case 5:
		fmt.Println("You can move to 5 spot")

	case 6:
		fmt.Println("You can move to 6 spot")

	default:
		fmt.Println("What was that!")
	}
}
