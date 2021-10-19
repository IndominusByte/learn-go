package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Nyoman Pradipta Dewantara", "Nyo"))
	fmt.Println(strings.Contains("Nyoman Pradipta Dewantara", "ayo"))
	fmt.Println(strings.Split("oman,pradipta,dewantara", ","))
	fmt.Println(strings.ToLower("NYOMAN PRADIPTA DEWANTARA"))
	fmt.Println(strings.ToUpper("nyoman pradipta dewantara"))
	fmt.Println(strings.Trim("            oman           ", " "))
	fmt.Println(strings.ReplaceAll("oman oman oman", "oman", "pradipta"))
}
