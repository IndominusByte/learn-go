package main

import "fmt"

type Address struct {
	Province, City, Subdisctric string
}

func changeProvinceToPapua(address *Address) {
	address.Province = "Papua"
}

func main() {
	var address1 Address = Address{"Bali", "Kuta Selatan", "Jimbaran"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.Province = "Balikpapan"

	*address2 = Address{"JKT", "Depok", "Holywings"}

	fmt.Println(address1)  // {JKT Depok Holywings}
	fmt.Println(*address2) // {JKT Depok Holywings}
	fmt.Println(*address3) // {JKT Depok Holywings}

	var address4 *Address = new(Address)
	// *address4 = Address{"lol", "lol", "lol"}

	fmt.Println(*address4)

	var alamat = Address{
		Province:    "",
		City:        "Ntt",
		Subdisctric: "sambuk",
	}

	changeProvinceToPapua(&alamat)

	fmt.Println(alamat)
}
