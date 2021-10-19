package main

import "fmt"

func main() {
	// style 1
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan")
	}

	fmt.Println()

	// style 2
	var i = 0

	for ; i < 3; i++ {
		fmt.Println("perulangan")
	}

	fmt.Println()

	// style 3
	for i := 0; ; i++ {
		if i == 3 {
			break
		}
		fmt.Println("perulangan")
	}

	fmt.Println()

	// style 4
	i = 0
	for i < 3 {
		fmt.Println("perulangan")
		i++
	}

	fmt.Println()

	var persons = [...]string{"oman", "pradipta", "dewantara", "indah"}

	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i])
	}

	for index, value := range persons {
		fmt.Println("index:", index, "=", value)
	}

	for _, value := range persons {
		fmt.Println(value)
	}

	for index := range persons {
		fmt.Println(index)
	}
}
