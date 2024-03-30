package main

import "fmt"

func main(){
    arrays := [...]int{
        122,
        124,
        457,
        124,
    }
    temp := arrays[2]
    fmt.Println("Panjang array adalah", len(arrays))
    fmt.Println("Value dari index ke - 2, adalah", arrays[2])
    arrays[2] = 100000
    fmt.Println("Rubah value dari nilai:", temp, "ke nilai:", arrays[2])
}