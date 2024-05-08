package main

import "fmt"

func main()  {
	fmt.Println("welcome to loops in go lang")
	days :=[]string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}
	for i:=0;i<len(days);i++{
		fmt.Println(days[i])
		if (days[i]=="Wednesday"){
                 break
		}
	}
	// for full loop
	for i:=range days{
		fmt.Println(days[i])
	}
	// for each loop 
	for index,day :=range days{
		fmt.Printf("The index is %v and day is %v\n",index,day)
	}
	// while loop
	rogueValue:=1
	for rogueValue <10{
		if rogueValue==2{
			goto lco
		}
       if rogueValue ==5 {
		break
	   }
		fmt.Println("the value is",rogueValue)
		rogueValue++
	}
	lco:
	     fmt.Println("learn at codeoneline.com")
}