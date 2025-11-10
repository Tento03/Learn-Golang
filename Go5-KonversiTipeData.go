package main

import "fmt"

func conversiondatatype() {
	var num1 = 20
	var num2 = float32(num1)
	var str1 = "Tento"
	var str2 = str1[1]
	var str3 = string(str2)

	fmt.Println("Num 1:", num1)
	fmt.Println("Num 2:", num2)
	fmt.Println("Str 1:", str1)
	fmt.Println("Str 2:", str2)
	fmt.Println("Str 3:", str3)
}
