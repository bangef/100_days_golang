package main

import (
	"fmt"
	"reflect"
)

func main(){
	var(
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)

		name string = "BAba"
		e = name[0]
		eString = string(e)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(name)
	// AKAN MENAMPILKAN TIPE DATA BYTE ATAU ALIAS DARI unit8
	fmt.Println(e, ", dengan tipe data:", reflect.TypeOf(e))
	fmt.Println(eString)
}