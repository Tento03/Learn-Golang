package main

import "fmt"

func mapping() {
	//Map Kosong
	var map0 map[string]int
	fmt.Println(map0)

	//Deklarasi Map
	var map1 = map[string]int{"tento": 1, "tendou": 2}
	fmt.Println(map1)

	//Make Map
	map2 := make(map[string]string)
	map2["Tendou"] = "kabuto"
	map2["Takumi"] = "Faiz"
	fmt.Println(map2["Tendou"])

	//Cek key
	key, value := map2["Tendou"]
	if value {
		fmt.Println("Key exists:", key)
	} else {
		fmt.Println("Key not exists")
	}

	//Hapus key
	delete(map2, "Tendou")
	fmt.Println(map2)

	//Looping map
	for k, v := range map1 {
		fmt.Println(k, ":", v)
	}
}
