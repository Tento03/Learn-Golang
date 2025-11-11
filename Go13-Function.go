package main

import "fmt"

//function tanpa parameter dan return
func basicfunction() {
	fmt.Println("Hello World")
}

//function dengan parameter
func funcParam(a int, b int) {
	var result = a + b
	fmt.Println("a + b :", result)
}

//function with return
func funRet1(a int, b int) int {
	var result = a * b
	return result
}

//function with multiple return
func funRet2(a int, b int) (int, int) {
	var min = a - b
	var div = a / b
	return min, div
}

//function with named return values
func funNamed(a int, b int) (resA int, resB int) {
	resA = a % b
	resB = b % a
	return
}

//function recursive
func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}
