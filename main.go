package main

import "fmt"

func main() {
	var p1 Person
	p1.Name = "Tento"
	p1.Age = 21
	p1.Adress.City = "Kotlin"
	p1.Adress.code = 1

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Adress.City)
	fmt.Println(p1.Adress.code)

	p2 := Person{
		Name: "Tendou",
		Age:  35,
	}
	fmt.Println(p2.Name)
	fmt.Println(p2.Age)

}
