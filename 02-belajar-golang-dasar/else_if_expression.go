package main

import "fmt"

func main() {
	// interface{} DIGUNAKAN UNTUK MEMBUAT TIPE DATA SECARA DINAMIS
	// NAMUN PERLU DIKONVERSI KE TIPE DATA SEBENARNYA
	mappedPerson := make(map[string]interface{})
	mappedPerson["name"] = "Bangef"
	mappedPerson["math value"] = -1
	mappedPerson["isPass"] = false

	// INI ADALAH CONTOH KONVERSI VALUE KE TIPE DATA INT
	if mathValue, ok := mappedPerson["math value"].(int); ok {
		if mathValue > 90 && mathValue <= 100 {
			mappedPerson["isPass"] = true
		} else if mathValue > 80 && mathValue <= 90 {
			mappedPerson["isPass"] = true
		} else if mathValue >= 0 && mathValue <= 80 {
			mappedPerson["isPass"] = false
		} else {
			fmt.Println("Error")
		}

		fmt.Println(mappedPerson)
	}
}
