package main

import "fmt"

func main() {
	var name = "kurniawan"

	if name == "oman" {
		fmt.Println("hai oman")
	} else if name == "joko" || name == "oman" {
		fmt.Println("hai", name)
	} else {
		fmt.Println("kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	}
}
