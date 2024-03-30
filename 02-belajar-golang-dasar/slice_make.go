package main

import "fmt"

func main() {
	// fmt.Println()
	sliceNames := make([]string, 2, 5)
	sliceNames[0] = "Budi"
	sliceNames[1] = "Eko"
	fmt.Println(sliceNames)
	fmt.Println("Panjang slice", len(sliceNames))
	fmt.Println("Kapasitas slice", cap(sliceNames))

	// KARENA KAPASITASNYA ADA 5 MAKA DAPAT MENAMBAHKAN DATA KE DALAM SLICE TANPA PERLU MEMBUAT ARRAY BARU LAGI
	sliceNames2 := append(sliceNames, "Caca")
	fmt.Println(sliceNames2)
	fmt.Println("Panjang slice", len(sliceNames2))
	fmt.Println("Kapasitas slice", cap(sliceNames2))

	// LALU COBA UBAH DATA DENGAN INDEX KE - 1
	sliceNames2[1] = "Rudi"
	fmt.Println(sliceNames2)
	// MAKA ARRAY REFERENSINYA JUGA IKUT BERUBAH KARENA DATA MASIH DIBAWAH KAPASITAS MAKSIMAL
	fmt.Println(sliceNames)
}
