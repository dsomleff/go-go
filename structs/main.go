package main

import "fmt"

type cat struct {
	firstName string
	lastName  string
}

func main() {
	var batat cat
	batat.firstName = "Batat"
	batat.lastName = "The Cat"

	fmt.Println(batat)
	fmt.Printf("%+v", batat)
}
