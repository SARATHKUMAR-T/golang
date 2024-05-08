package main

import "fmt"

func main()  {
	fmt.Println("welcome to ifelse section")

	loginCount := 4

	var result string

	if loginCount < 10 {
		result = "regular user"
	}else{
		result = "something else"
	}
	fmt.Println(result)

	if num :=6; num <10{
		fmt.Println("number is less then 10")
	}else{
		fmt.Println("Number is greater then 10")
	}
}