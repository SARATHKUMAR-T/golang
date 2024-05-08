package main

import "fmt"


func main()  {
	fmt.Println("welcome to methods section")

	// no inheritance in go. and no Super or Parent.

	victor := User{"NirmalVictor","victor@gmail.com",true,24}
        fmt.Println(victor)
		fmt.Printf("victor details are %+v\n",victor)
		fmt.Printf("name is : %v\n",victor.Name)
		victor.GetStatus()

}

	type User struct {
		Name string
		Email string
		Status bool
		Age int
	}
	// methods n go lang
	func (u User) GetStatus()  {
		fmt.Println("the status is",u.Status)
	}