package main

import "fmt"

func main() {
	var arrayCharacter = [...]string{"a", "b", "c", "d"}
	var sliceCharacter []string = arrayCharacter[:3] // [0:3]
	var lengthArray int = len(arrayCharacter)
	var lengthSlice int = len(sliceCharacter) // 3 - 0 = 3 => POINTER AKHIR - POINTER AWAL
	fmt.Println("Panjang dari arrayCharacter adalah", lengthArray)
	fmt.Println("Panjang dari sliceCharacter adalah", lengthSlice)
}
