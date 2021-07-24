package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	zipCode int
}

func main() {
	var alex person

	// hisam := person{firstName: "hisam", lastName: "Maulana"}
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	// nuri := person{
	// 	firstName: "Nuri",
	// 	lastName: "",
	// 	contact: contact{
	// 		email: "nuri@hisamcode.id",
	// 		zipCode: 25000,
	// 	},
	// }

	// hisam.print()
	// nuri.print()

	alexPointer := &alex
	fmt.Println(alexPointer)
	alexPointer.updateName("alesss")
	// alex.updateName("aa")
	alex.print()

}

func (p person) print () {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person ) updateName (newFirstName string) {
	(pointerToPerson.firstName) = newFirstName
}