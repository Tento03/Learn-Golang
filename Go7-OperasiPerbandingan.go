package main

import "fmt"

func opcompare() {
	var a = 2
	var b = 3
	var result1 = a == b
	var result2 = a >= b
	var result3 = a <= b
	var result4 = a != b

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
