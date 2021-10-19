package main

import "fmt"

func Upss(i int) interface{} {
	switch i {
	case 1:
		return 1
	case 2:
		return true
	default:
		return "Upss"
	}
}

func main() {
	var kosong interface{} = Upss(3)
	fmt.Println(kosong)
}
