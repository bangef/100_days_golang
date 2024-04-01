package main

import "fmt"

func main() {
	// MEMBUAT MAP DENGAN 2 VERSI:
	// VERSI 1 :
	mappedPerson := make(map[string]string)
	mappedPerson["name"] = "Bangef"
	mappedPerson["hobby"] = "Learning Code"

	fmt.Println(mappedPerson)
	fmt.Println(mappedPerson["name"])
	fmt.Println(mappedPerson["hobby"])
	// VERSI 2 :
	mappedPerson2 := map[string]string{}
	mappedPerson2["name"] = "Jhone Doe"
	mappedPerson2["hobby"] = "Swimming"

	fmt.Println(mappedPerson2)
	fmt.Println(mappedPerson2["name"])
	fmt.Println(mappedPerson2["hobby"])

	// DAPAT DIGUNAKAN UNTUK MENGHAPUS VALUE
	// NAMUN JIKA INGIN MENGAKSES KEYNYA MASIH DAPAT DILAKUKAN
	// DAN TIDAK ERROR
	delete(mappedPerson, "hobby")
	fmt.Println(mappedPerson)
	fmt.Println(mappedPerson["hobby"])
}
