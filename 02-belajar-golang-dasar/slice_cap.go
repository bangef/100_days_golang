package main

import "fmt"

func main() {
	arrayNumbers := [...]int{
		1,
		45,
		21,
		3253,
		2,
		965,
		965,
		965,
		965,
		965,
		965,
		965,
		965,
		965,
		965,
		965,
	}
	sliceNumbers := arrayNumbers[3:6]
	capacitySliceNumbers := cap(sliceNumbers) // 16 - 3 = 13 => PANJANG ARRAY - POINTER AWAL
	fmt.Println("Kapasitas sliceNumbers adalah", capacitySliceNumbers)
}
