package main

import "fmt"

func arrays() {
	//Deklarasi Array
	var array1 [3]int
	fmt.Println(array1)

	var riders = [3]string{"kiva", "build", "kabuto"}
	fmt.Println(riders)

	autoArray := []int{10, 20, 30}
	fmt.Println(autoArray)

	//Akses Array
	fmt.Println("Rider index 0:", riders[0])

	//For Biasa
	for i := 0; i < len(riders); i++ {
		fmt.Println("Rider ke ", i, " : ", riders[i])
	}

	//For Range
	for index, value := range riders {
		fmt.Println(index, value)
	}
}
