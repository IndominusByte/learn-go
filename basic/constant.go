package main

import "fmt"

func main() {
	// tidak error jika unused variable
	const firstName string = "oman"
	const lastName = "dewantara"
	const age = 10

	const (
		nim   uint32 = 170010057
		phone uint64 = 87862265363
	)

	fmt.Println(age)
}
