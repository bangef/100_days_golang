package main

import "fmt"

func main(){
	var (
		name1 = "Bangef"
		name2 = "Bangef"
		number1 = 1
		number2 = 2
		result bool
	)

	result = (name1 == name2)
	fmt.Println("Apakah",name1,"sama dengan",name2,result)
	result = (name1 != name2)
	fmt.Println("Apakah",name1,"tidak sama dengan",name2,result)
	result = (number1 > number2)
	fmt.Println("Apakah",number1,"lebih dari",number2,result)
	result = (number1 < number2)
	fmt.Println("Apakah",number1,"kurang dari",number2,result)
	result = (number1 >= number2)
	fmt.Println("Apakah",number1,"lebih dari sama dengan",number2,result)
	result = (number1 <= number2)
	fmt.Println("Apakah",number1,"kurang dari sama dengan",number2,result)
}