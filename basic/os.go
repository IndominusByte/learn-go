package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for index, value := range args {
		fmt.Println(index, value)
	}

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	}

	bath_theme := os.Getenv("BAT_THEME")
	fmt.Println(bath_theme)
}
