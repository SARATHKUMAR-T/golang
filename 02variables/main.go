package main

import "fmt"

const LoginToken = "dwudfbwfbewfbef"

func main()  {
	var userName string = "Sarath"
	fmt.Println(userName)
	fmt.Printf("Variable is a type: %T \n",userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is a type: %T \n",isLoggedIn)

	var anotherVariable int
	fmt.Println((anotherVariable))
	fmt.Printf("Variable is a type: %T \n",anotherVariable)

	var anotherVariable2 string
	fmt.Println((anotherVariable))
	fmt.Printf("Variable is a type: %T \n",anotherVariable2)

	var anotherVariable3 bool
	fmt.Println((anotherVariable))
	fmt.Printf("Variable is a type: %T \n",anotherVariable3)


	fmt.Printf("Variable is a type: %T \n",LoginToken)

	// implicit type

	var website = "www.google.com"
	fmt.Println(website)

	// no var style

	numberOfUsers :=12000
	fmt.Println(numberOfUsers)
	numberOfUsers=45
}