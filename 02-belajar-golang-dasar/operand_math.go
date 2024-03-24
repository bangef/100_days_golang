package main

import (
	"fmt"
	"math/rand" // PAKET UNTUK MENDAPATKAN RANDOM NUMBER 
)

func main(){
	var(
		a int = rand.Intn(50) // UNTUK GENERATE RANDOM INT
		b int = rand.Intn(50)
		c int = rand.Intn(50)
		d int = rand.Intn(50)
		result int
	)

	fmt.Println("All Data: a =>",a, ", b =>", b, ", c =>", c, ", d =>", d)
	result = a + b * c - d
	fmt.Println("Formula: a + b * c - d, Hasilnya:", result)
}