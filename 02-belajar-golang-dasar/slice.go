package main

import "fmt"

func main() {
	arrayString := [12]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	// array[low:high]
	sliceString := arrayString[0:2]
	fmt.Println("Nilai dari slice:", sliceString)
	fmt.Println("Panjang dari slice:", len(sliceString))   // 2 - 0 = 2
	fmt.Println("Kapasitas dari slice:", cap(sliceString)) // 12 - 0 = 12
	fmt.Println("--------")
	// array[low:]
	sliceString2 := arrayString[5:] // SAMA DENGAN array[5:12]
	fmt.Println("Nilai dari slice:", sliceString2)
	fmt.Println("Panjang dari slice:", len(sliceString2))   // 12 - 5 = 7
	fmt.Println("Kapasitas dari slice:", cap(sliceString2)) // 12 - 5 = 7
	fmt.Println("--------")
	// array[:high]
	sliceString3 := arrayString[:4] // SAMA DENGAN array[0:4]
	fmt.Println("Nilai dari slice:", sliceString3)
	fmt.Println("Panjang dari slice:", len(sliceString3))   // 4 - 0 = 4
	fmt.Println("Kapasitas dari slice:", cap(sliceString3)) // 12 - 0 = 12
	fmt.Println("--------")
	// array[:]
	sliceString4 := arrayString[:] // SAMA DENGAN array[0:12]
	fmt.Println("Nilai dari slice:", sliceString4)
	fmt.Println("Panjang dari slice:", len(sliceString4))   // 12 - 0 = 12
	fmt.Println("Kapasitas dari slice:", cap(sliceString4)) // 12 - 0 = 12
	fmt.Println("--------")
	// ERROR
	// sliceString5 := arrayString[5:1] // POINTER AWAL TIDAK BOLEH LEBIH BESAR DARI POINTER AKHIR
	// fmt.Println("Nilai dari slice:", sliceString5)
	// fmt.Println("Panjang dari slice:", len(sliceString5))   // 12 - 0 = 12
	// fmt.Println("Kapasitas dari slice:", cap(sliceString5)) // 12 - 0 = 12

	// KESIMPULAN:
	// 1. APABILA SLICE TIDAK MEMBERITAHU POINTER AKHIRNYA MAKA BISA DIPASTIKAN NILAI LEN DAN CAPACITYA SAMA
	// 2. NAMUN APABILA TIDAK MENENTUKAN NILAI POINTER AWALNYA, CONTOH `arrayString[:4]`. MAKA NILAI CAPACITYNYA SAMA DENGAN PANJANG ARRAY REFERENSINYA.
	// 3. APABILA POINTER AWAL DAN AKHIRNYA DI DEKLARASI, CONTOH `arrayString[5:100]`. MAKA RUMUSNYA:
	// 		LEN = POINTER AKHIR - POINTER AWAL
	// 		CAP = PANJANG ARRAY - POINTER AWAL
}
