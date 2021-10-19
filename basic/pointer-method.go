package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	name := Man{"oman"}
	name.Married()

	fmt.Println(name)
}
