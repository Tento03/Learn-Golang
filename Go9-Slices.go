package main

import "fmt"

func slice() {
	//Inisiasi Slice
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	//Slice dari Array
	arr1 := []int{2, 4, 6}
	slice2 := arr1
	fmt.Println(slice2)

	//Slice dari make
	slice3 := make([]int, 3)
	fmt.Println(slice3)

	//Tambah Elemen
	slice4 := []int{1, 2, 3}
	slice4 = append(slice4, 4, 5, slice2[0])
	fmt.Println(slice4)

	//Ubah Elemen
	slice4[5] = 6
	fmt.Println(slice4)

	//Hapus Elemen
	slice5 := []int{1, 2, 3, 4, 5}
	slice5 = append(slice5[:2], slice5[3:]...)
	fmt.Println(slice5)

	//For Biasa
	for i := 0; i < len(slice5); i++ {
		fmt.Println("Index ke ", i, ":", slice5[i])
	}

	//For Range
	for index, value := range slice5 {
		fmt.Println("Index ke ", index, ":", value)
	}
}
