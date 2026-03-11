package main

import "fmt"

type cat struct {
	firstName string
	lastName  string
	contact   familyContact
}

type familyContact struct {
	contactName  string
	contactPhone int
}

func main() {
	batat := cat{
		firstName: "Batat",
		lastName:  "The Cat",
		contact: familyContact{
			contactName:  "Ju",
			contactPhone: 123,
		},
	}

	fmt.Println(batat)
	fmt.Printf("%+v", batat)
}
