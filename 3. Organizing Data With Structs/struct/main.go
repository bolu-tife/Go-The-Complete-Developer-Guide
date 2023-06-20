package main

import "fmt"

type contactInfo struct {
	zipCode int
	email   string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //(*pointerToPerson) - is the value
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@alex.com",
			zipCode: 100100,
		}}

	// alexPointer := &alex
	// alexPointer.updateName("Noe")

	alex.updateName("Noe") //this works because Go automatically picks out the memory(pointer) because  of the receiver type is expecting a pointer - pointerToPerson *person
	alex.print()
}
