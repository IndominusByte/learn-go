package main

import (
	"example/src/basic/helper"
	"fmt"
)

func main() {
	helper.SayHello("oman")
	// helper.sayGoodBye("pradipta") error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) error
}
