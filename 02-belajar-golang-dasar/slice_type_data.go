package main

import "fmt"

func main() {
	arrayString := [124]string{"Budi", "Joko", "Prabowo", "Kaesang", "BJ. Habibi"}
	// MEMBUAT SLICE
	// POINTER AWAL		= 45
	// POINTER AKHIR	= 99
	// LENGTH					= 99 - 45 	= 54 // POINTER AKHIR - POINTER AWAL
	// CAPACITY				= 124 - 45 	= 79 // PANJANG REFERENSI ARRAY - POINTER AWAL
	slice1 := arrayString[45:99]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
