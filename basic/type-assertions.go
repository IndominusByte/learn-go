package main

import "fmt"

func random() interface{} {
	return "oman"
}

func main() {
	rand := random()

	fmt.Println(rand.(string))

	switch value := rand.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type data")
	}
}
