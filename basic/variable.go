package main

import "fmt"

func main() {
	// style 1
	var name string

	name = "Oman"
	fmt.Println(name)

	name = "Pradipta"
	fmt.Println(name)

	// style 2
	var friend = "Doggy"
	fmt.Println(friend)

	// style 3
	benefit := "Married"
	fmt.Println(benefit)

	// style 4
	var (
		firstname = "oman"
		lastname  = "dewantara"
	)
	fmt.Println(firstname + lastname)

	// style 5
	first, last := "oman", "dewantara"
	fmt.Println(first, last)

	// style 6
	var a, b int
	a = 1
	b = 2
	fmt.Println(a, b)
}
