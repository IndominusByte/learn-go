package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put you host database")
	username := flag.String("username", "root", "Put you username database")
	password := flag.String("password", "root", "Put you password database")
	port := flag.Int("port", 80, "Put you port application")

	flag.Parse()

	fmt.Println(*host, *username, *password, *port)
}
