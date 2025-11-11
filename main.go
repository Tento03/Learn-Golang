package main

import "fmt"

func main() {
	basicfunction()
	funcParam(2, 3)

	result := funRet1(2, 3)
	fmt.Println("a * b:", result)

	result2, result3 := funRet2(4, 6)
	fmt.Println("a - b:", result2)
	fmt.Println("a / b:", result3)

	result4, result5 := funNamed(2, 3)
	fmt.Println("a % b:", result4)
	fmt.Println("b % a:", result5)

	resultRecursive := recursive(4)
	fmt.Println(resultRecursive)
}
