package main

import "fmt"

func condition() {
	var traffic = map[int]string{0: "merah", 1: "kuning", 2: "hijau"}
	var con string

	for i := 0; i <= len(traffic); i++ {
		if traffic[i] == "merah" {
			con = "Stop"
			fmt.Println(con)
		} else if traffic[i] == "kuning" {
			con = "Ready"
			fmt.Println(con)
		} else if traffic[i] == "hijau" {
			con = "Go"
			fmt.Println(con)
		}
	}

}
