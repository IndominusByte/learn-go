package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "nyoman"
	names[1] = "pradipta"
	names[2] = "dewantara"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)

	x := [3]int{10, 20, 30}
	y := [...]int{1, 2, 3, 4, 5}

	fmt.Println(x, y)
}
