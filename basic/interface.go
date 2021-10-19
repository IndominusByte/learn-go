package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Kind string
}

func (animal Animal) getName() string {
	return animal.Kind
}

func main() {
	person := Person{"oman"}

	sayHello(person)

	animal := Animal{
		Kind: "mamalia",
	}

	sayHello(animal)
}
