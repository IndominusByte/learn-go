package main

import "fmt"

func main() {
	var i = 10
	i += 10
	fmt.Println(i) // 20

	var a = 2
	var b = 3
	var c = a + b
	fmt.Println(c) // 5

	c++
	fmt.Println(c) // 6

	var negative = -100
	var positive = +100
	fmt.Println(negative, positive)
}
