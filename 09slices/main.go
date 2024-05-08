package main

import (
	"fmt"
	"sort"
)


func main()  {
	fmt.Println("welcome to slices section")

	var fruitList =[]string{"apple","tomato","peach"}
	fmt.Printf("type of fruitlist is %T \n",fruitList)

	fruitList = append(fruitList,"banana","grapes" )
	// fmt.Println("fruitlist is",fruitList)
	// newList :=append(fruitList[1:3])
	// fmt.Println("newList",newList)
	// fmt.Println("fruitList",fruitList)
	highScore :=make([]int,4)
	highScore[0]=432
	highScore[1]=365
	highScore[2]=680
	highScore[3]=836
	highScore=append(highScore,555,120,247 )
	sort.Ints(highScore)
    // fmt.Println("highscores",highScore)
    // fmt.Println(sort.IntsAreSorted(highScore))
	var courses =[]string{"reactjs","python","javascript","ruby"}
	var index int =2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println("courses",courses)
}