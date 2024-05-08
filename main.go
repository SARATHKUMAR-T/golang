package main

import (
	"fmt"
	"io"
	"os"
)


func main()  {
	fmt.Println("welcome to files section")

	content := "these needs to go in a file."

	file,err:=os.Create("./logfile.txt")
	checkNillErr(err)
	length,err:=io.WriteString(file,content)
	checkNillErr(err)
	fmt.Println("the length is",length)
	defer file.Close()
    readFile("./logfile.txt")
}


func readFile(fileName string){
	dataByte,err:=os.ReadFile(fileName)
	checkNillErr(err)
	fmt.Println("the databyte is \n",string(dataByte))
}

func checkNillErr(err error){
if err != nil {
		panic(err)
	}
}