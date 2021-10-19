package main

import "fmt"

type Person struct {
	Name, State string
	age         int
}

func (person Person) greetingTo(name string) {
	fmt.Println("Hi my name is", person.Name, "you are", name, "?")
}

func (person Person) myAge() {
	fmt.Println("my age is", person.age)
}

func main() {
	var person_1 Person
	person_1.Name = "oman"
	person_1.State = "indonesia"
	person_1.age = 30

	fmt.Println(person_1)

	person_1.greetingTo("dewantara")
	person_1.myAge()

	person_2 := Person{
		Name:  "oman",
		State: "indonesia",
		age:   22,
	}

	fmt.Println(person_2)

	person_3 := Person{"pradipta", "bali", 21}

	fmt.Println(person_3)

	person_3.greetingTo("dewantara")
	person_3.myAge()
}
