package helper

import "fmt"

const Application = "belajar golang"
const version = "v1.0.0"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Goodbye", name)
}
