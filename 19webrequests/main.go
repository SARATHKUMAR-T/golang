package main

import (
	"fmt"
	"io"
	"net/http"
)

const url="https://lco.dev"

func main()  {
	fmt.Println("welcome to web request")

	response,err:=http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the response type is %T \n",response)

	defer response.Body.Close()     //callers responsibility to close the connection.

	databytes,err:= io.ReadAll(response.Body)

	fmt.Println("data",databytes)
}
