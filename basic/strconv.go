package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	num, err := strconv.ParseInt("10000", 10, 64)
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err.Error())
	}
}
