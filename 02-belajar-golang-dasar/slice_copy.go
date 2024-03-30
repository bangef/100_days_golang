package main

import "fmt"

func main() {
	arrayNumbers := [5]int{
		1,
		45,
	}

	sliceOne := arrayNumbers[:]
	sliceTwo := make([]int, len(sliceOne), cap(sliceOne))

	// MENYALIN NILAI DARI "sliceOne" KE "sliceTwo"
	copy(sliceTwo, sliceOne)

	// MENAMPILKAN NILAI "sliceOne" dan "sliceTwo"
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)
}
