package main

import "fmt"

func main(){
    var (
        number1    = 10
        number2    = 20
        result bool
    )
    
    // APAKAH OUTPUT JIKA TRUE
    // HASILNYA ADALAH FALSE
    result    = (number1 < number2)
    fmt.Println("Hasilnya adalah", !result)
    // APAKAH OUTPUT JIKA FALSE
    // HASILNYA ADALAH TRUE
    result    = (number1 > number2)
    fmt.Println("Hasilnya adalah", !result)
}