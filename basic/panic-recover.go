package main

import "fmt"

func runApp(err bool) {
	defer func() {
		if message := recover(); message != nil {
			fmt.Println("Error occur", message)
		}
	}()

	if err {
		panic("Error bro")
	}

	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("aplikasi berlanjut")
}
