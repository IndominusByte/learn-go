package main

import "fmt"

func main() {
	var name = "joko"

	switch name {
	case "oman", "joko":
		fmt.Println("hai oman")
		fmt.Println("hai oman")
	case "eko", "kurniawan":
		fmt.Println("hai eko")
		fmt.Println("hai eko")
	default:
		fmt.Println("kenalan donk")
		fmt.Println("kenalan donk")
	}

	/*
		switch length := len(name); length > 5 {
		case true:
			fmt.Println("terlalu panjang")
		case false:
			fmt.Println("terlalu pendek")
		}
	*/

	length := len(name)

	switch {
	case length > 5 || length == 4:
		fmt.Println("lebih dari 5")
	case length < 5:
		fmt.Println("kurang dari 5")
	}
}
