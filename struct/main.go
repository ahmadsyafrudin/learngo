package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "john@me.local",
			zipCode: 9400,
		},
	}

	fmt.Printf("%+v", john)
}
