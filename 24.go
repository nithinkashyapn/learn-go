package main

import (
	"fmt"
)

type Contact struct {
	name  string
	phone uint64
	email string
}

func display(person Contact) {
	fmt.Printf("Name is %v\nPhone is %v\nemail is %v\n", person.name, person.phone, person.email)
}

func main() {

	var person1 Contact

	person1.name = "Foo"
	person1.phone = 1111111111
	person1.email = "foo@example.com"

	person2 := Contact{name: "Bar", phone: 2222222222, email: "bar@example.com"}

	display(person1)
	display(person2)

}
