package main

import "fmt"

type cat struct {
	firstName string
	lastName  string
	familyContact
}

type familyContact struct {
	contactName  string
	contactPhone int
}

func main() {
	batat := cat{
		firstName: "Batat",
		lastName:  "The Cat",
		familyContact: familyContact{
			contactName:  "Ju",
			contactPhone: 123,
		},
	}

	batat.print()
	batat.updateName("Batya")
	batat.print()
}

func (c cat) print() {
	fmt.Println(c)
	fmt.Printf("%+v", c)
}

func (c *cat) updateName(newFirstName string) {
	(*c).firstName = newFirstName
}
