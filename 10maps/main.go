package main

import "fmt"

func main()  {
	fmt.Println("welcome to map section")

	laungauge := make(map[string]string)

	laungauge["JS"] = "javascript"
	laungauge["PY"] = "python"
	laungauge["RB"] = "Ruby"

	fmt.Println("map",laungauge)
    fmt.Println("the js stands for",laungauge["JS"])

	delete(laungauge,"RB")

	fmt.Println(laungauge)

	for key,value := range laungauge {
		fmt.Printf("For key of %v,the value is %v\n",key,value)
	}
}