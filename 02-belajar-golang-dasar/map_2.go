package main

import "fmt"

func main() {
	mappedPerson := map[string]string{
		"name":  "Bangef",
		"hobby": "Learning Code",
	}

	// MENAMPILKAN NILAI VALUE DARI KEY "name"
	fmt.Println(mappedPerson["name"])
	// MENAMPILKAN NILAI VALUE DARI KEY "hobby"
	fmt.Println(mappedPerson["hobby"])
}
