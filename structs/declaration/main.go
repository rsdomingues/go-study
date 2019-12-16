package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// it is not necessary to define the name of the field.
// if you inform only the type, the name of the field will be the name of the struct.
type person struct {
	firsname string
	lastname string
	contact  contactInfo
}

func main() {
	john := person{"John", "Smith", contactInfo{"john@smith.com", 12345}}
	jack := person{firsname: "Jack", lastname: "Anderson", contact: contactInfo{email: "jack@anderson.com", zipCode: 12345}}
	var alex person

	alex.firsname = "Alex"
	alex.lastname = "Hardy"
	alex.contact.email = "alex@hardy.com"
	alex.contact.zipCode = 12345

	fmt.Printf("%+v\n", john)
	fmt.Printf("%+v\n", jack)
	fmt.Printf("%+v\n", alex)
}
