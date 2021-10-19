package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	name := NewMap("oman")
	if name == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(name)
	}
}
