package main

import "fmt"

func logging() {
	fmt.Println("function selesai dijalankan")
}

func runApplication(val int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / val
	fmt.Println("result", result)
}

func main() {
	runApplication(0)
}
