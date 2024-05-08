package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main()  {
	fmt.Println("Switch and case in go lang")
	rand.Seed(time.Now().UnixNano())
	diceNumber :=rand.Intn(6)+1
	fmt.Println("the dice value is",diceNumber)

	switch diceNumber {
	case 1 :
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")	
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spots")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spots and open")
	default:
		fmt.Println("What was that!")					
	}
}