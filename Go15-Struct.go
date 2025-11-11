package main

type Person struct {
	Name   string
	Age    int
	Adress Adress
}

type Adress struct {
	City string
	code int
}
