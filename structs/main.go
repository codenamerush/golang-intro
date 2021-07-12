package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex := person{"Alex", "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.firstName = "Anderson"
	fmt.Println(alex)
}
