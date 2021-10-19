package main

import (
	"errors"
	"fmt"
)

func Pembagi(value int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak bisa 0")
	} else {
		return value / pembagi, nil
	}
}

func main() {
	result, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("error", err.Error())
	}
}
