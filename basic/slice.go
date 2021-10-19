package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari", "Februari",
		"Maret", "April",
		"Mei", "Juni",
		"Juli", "Agustus",
		"September", "Oktober",
		"November", "Desember",
	}

	var slice = months[4:7]

	fmt.Println(slice)      // [Mei Juni Juli]
	fmt.Println(len(slice)) // 3
	fmt.Println(cap(slice)) // 8

	months[5] = "June"
	fmt.Println(slice) // [Mei June Juli]

	slice[0] = "May"
	fmt.Println(months) // [Januari Februari Maret April (May) <-- June Juli Agustus September Oktober November Desember]

	var slice2 = months[11:]
	fmt.Println(slice2) // [Desember]

	var slice3 = append(slice2, "oman")
	fmt.Println(slice3) // [Desember oman]
	slice3[0] = "Bukan Desember"
	fmt.Println(slice3) // [Bukan Desember oman]
	fmt.Println(slice2) // [Desember]

	// MAKE SLICE
	newSlice := make([]string, 2, 5)

	newSlice[0] = "nyoman"
	newSlice[1] = "pradipta"

	fmt.Println(newSlice)      // [nyoman pradipta]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) // [nyoman pradipta]

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray, iniSlice) // [1 2 3 4 5] [1 2 3 4 5]
}
