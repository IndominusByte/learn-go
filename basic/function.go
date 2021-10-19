package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func myName(name string) string {
	return "my name is " + name
}

func getMyName() (string, string) {
	return "Nyoman", "Pradipta"
}

// named return function
func getFullName() (firstName, middleName, lastName string) {
	firstName, middleName, lastName = "nyoman", "pradipta", "dewantara"
	return
}

// variadic function
func sumAll(numbers ...int) (total int) {
	total = 0

	for _, value := range numbers {
		total += value
	}

	return
}

// function value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// function as parameter

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filtered := filter(name)
	fmt.Println("Hello", filtered)
}

func filterSpam(name string) string {
	switch name {
	case "anjing":
		return "..."
	default:
		return name
	}
}

// anonymous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursive function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	sayHello()
	sayHelloTo("nyoman", "dewantara")
	fmt.Println(myName("oman"))
	firstName, lastName := getMyName()
	fmt.Println(firstName, lastName)
	fmt.Println(getFullName())

	total := sumAll(10, 20, 30, 40)
	fmt.Println(total)

	var data = []int{1, 2, 3, 4, 5}
	total = sumAll(data...)
	fmt.Println(total)

	goodbye := getGoodBye
	result := goodbye("oman")
	fmt.Println(result)

	fmt.Println(getGoodBye("oman"))

	sayHelloWithFilter("oman", filterSpam)
	sayHelloWithFilter("anjing", filterSpam)

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("oman", blacklist)
	registerUser("admin", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))
}
