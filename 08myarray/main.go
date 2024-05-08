package main

import "fmt"


func main()  {
	fmt.Println("welcome to array section")

	var fruitList [4] string

	fruitList[0]="fig"
	fruitList[1]="grapes"
	fruitList[3]="apple"
   
	fmt.Println("values of fruitlist is :",fruitList)
	fmt.Println("values of fruitlist is :",len(fruitList))

	var vegList =[5] string {"mushroom","onion","potato"}
	fmt.Println("veglist is :",vegList)
	fmt.Println("veglist is :",len(vegList))
}