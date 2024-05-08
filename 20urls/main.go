package main

import (
	"fmt"
	"net/url"
)

const myrul string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=3e3rcxbd98sxc"

func main() {
	fmt.Println("welcome to url section in golang")
	fmt.Println(myrul)
	// parsing
	result, _ := url.Parse(myrul)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()

	for _, val := range qParams {
		fmt.Println(val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "localhost:3000",
		Path:    "/product",
		RawPath: "colour=black",
	}
	consturctedUrl := partsOfUrl.String()
	fmt.Println(consturctedUrl)
}
