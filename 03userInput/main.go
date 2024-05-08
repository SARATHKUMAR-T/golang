package main

import (
	"bufio"
	"fmt"
	"os"
)



func main (){
  welcome := "welcome to user input"
  fmt.Println(welcome)
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter the rating forour pizza:")

   //   comma ok syntax
    input,_ :=reader.ReadString('\n')
	fmt.Println("input",input)
	fmt.Printf("type of the input %T ",input)

}