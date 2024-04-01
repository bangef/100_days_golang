package main

import "fmt"

func main() {
	mappedPerson := map[string]string{
		"name":  "Bangef",
		"hobby": "Learning Code",
	}
	fmt.Println("Hello, my name", mappedPerson["name"], "and my hobby is", mappedPerson["hobby"])
	// DAPAT MENAMBAHKAN DENGAN CARA INI
	mappedPerson["salah"] = "test"
	fmt.Println(mappedPerson["salah"])

	// ATAUPUN APABILA KEY TIDAK DIDEKLARASI DAN TIDAK DI INISIAILISASI
	// MAKA DEFAULTNYA ADALAH STRING KOSONG
	fmt.Println(mappedPerson["tidak terdaftar"])

	// NAMUN AKAN TERJADI ERROR APABILA VALUE TIDAK SESUAI DENGAN TIPE DATANYA
	// ERROR : cannot use 20 (untyped int constant) as string value in assignment
	// mappedPerson["testing"] = 20
	// fmt.Println(mappedPerson["testing"])

	// DAN APABILA INGIN MENGUPDATE VALUE DARI KEY MAKA DAPAT ME RE-ASSIGN
	mappedPerson["salah"] = "iya salah"
	fmt.Println(mappedPerson["salah"])
	fmt.Println(mappedPerson)

}
