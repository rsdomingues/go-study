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
	person := person{firsname: "Jack", lastname: "Anderson", contact: contactInfo{email: "jack@anderson.com", zipCode: 12345}}

	// check function declaration comments to understand the diference between passing by value and pass by diagram.
	person.updateName("Alexander")
	person.pinterUpdateName("Alexandrus")

	person.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//update name will receive a copy value from the person object
//so the copy will have the name updated but the original value will remain the same
func (p person) updateName(firstname string) {
	p.firsname = firstname
}

// update name will receive now a pointer to the person object
// so when we alter the name, whe alter the actual original object
func (p *person) pinterUpdateName(firstname string) {
	// I can use p or *p, go will help you with by using it for you in this case.
	p.firsname = firstname
}
