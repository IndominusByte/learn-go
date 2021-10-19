package main

import "fmt"

func main() {
	// jika memaksakan convert dari besar ke kecil akan otomatis angka berubah menjadi max penampung
	// sebaliknya jika int min penampung
	var nilaiUint16 uint16 = 65535      // 65535
	var nilaiUint8 = uint8(nilaiUint16) // 255

	fmt.Println(nilaiUint16, nilaiUint8)

	var Int16 int16 = 130       // 130
	var Int8 int8 = int8(Int16) // -126

	fmt.Println(Int16, Int8)
}
