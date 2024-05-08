package main

import "fmt"

func main()  {
	defer fmt.Println("differ is awesome")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("welcome to differ section")
	myDefer()
}

func myDefer()  {
	for i:=0;i<5;i++{
	defer fmt.Println("\n",i)
	}
}