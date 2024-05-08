package main

import "fmt"


func main()  {
	fmt.Println("welcome to function section")
	result:=adder(45,5)
	fmt.Println("the result is :",result)
	prores,msg:=proAdder(2,3,4,5,8,9,4,443)
	fmt.Println("pro result is ",prores)
	fmt.Println("msg",msg)
	greeter()
}

func greeter() {
	fmt.Println("Vanakkam")
}

func adder(valOne int,valTwo int)int{
	return valOne+valTwo
}

func proAdder(values ...int)(int,string){
	total:=0
	for _,val :=range values{
		total+=val
	}
	return total,"hi there"
}