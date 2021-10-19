package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "oman",
		"address": "purigading",
	}

	person["title"] = "programmer"

	fmt.Println(person)          // map[address:purigading name:oman title:programmer]
	fmt.Println(person["name"])  // oman
	fmt.Println(person["title"]) // programmer

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "oman"
	book["wrong"] = "ups"

	fmt.Println(book) // map[author:oman title:belajar golang wrong:ups]

	delete(book, "wrong")

	fmt.Println(book) // map[author:oman title:belajar golang]
}
