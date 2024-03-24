package main

import "fmt"

func main(){

	// REFACTORING UNTUK MENGGANTIKAN KATA KUNCI VAR DENGAN ":="
	// DAPAT DILAKUKAN HANYA SAAT PENDEKLARASIAN DAN INISILISASI VRAIBLE
	name := "Bangef"
	age := 20
	height := 170

	fmt.Println(name)
	name = "Ficri Hanip"
	fmt.Println(name)
	fmt.Println("Umur saya: ", age)
	fmt.Println("Tinggi saya: ", height)
}