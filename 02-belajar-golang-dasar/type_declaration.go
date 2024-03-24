package main

import (
	"fmt"
	"reflect"
)

// INI ADALAH ALIAS ATAU NAMA LAIN DARI TIPE DATA STRING
type NoKtp string

func main(){
	var ktpBangef NoKtp = "111111111"

	fmt.Println(ktpBangef, "tipe data:", reflect.TypeOf(ktpBangef))
	// KONVERSI DARI TIPE DATA STRING KE NoKtp
	fmt.Println(NoKtp("22222222"), "tipe data:", reflect.TypeOf(NoKtp("22222222")));
}