package main

import (
	"fmt"
	"time"
)


func main()  {
	fmt.Println("Welcome to time section")

	presentTime := time.Now()
	fmt.Println("present time",presentTime)

	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createDate := time.Date(1998,time.November,01,22,10,0,0,time.Local)

	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println(createDate,"created Date")
}