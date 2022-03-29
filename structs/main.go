package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 2076,
		},
	}

	fmt.Println(alex)
	alex.print()

	// change alex's name to jimmy
	(&alex).updateName("jimmy")

	fmt.Println(alex)
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
