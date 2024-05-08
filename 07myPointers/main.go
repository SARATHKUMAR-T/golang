package main

import "fmt"


func main()  {
	fmt.Println("welcome to pointers section")

	// creating a pointer
	// var ptr *int 
	// fmt.Println("value of ptr",ptr)


	// referencing a pointer
	//  & means reference

	myNumber := 24
	var ptr= &myNumber
	fmt.Println("the value of ptr is",ptr)
	fmt.Println("the value of ptr is",*ptr)

	*ptr=*ptr + 6
	fmt.Println("the new value is",*ptr)
}