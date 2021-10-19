package main

import "fmt"

func main() {
	// make alias
	type Phone = string

	var omanPhone Phone
	omanPhone = "08782727277272"

	fmt.Println(omanPhone)
}
