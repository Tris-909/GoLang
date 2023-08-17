package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	phoneNumber string
	countryCode string
}

func main() {
	tri := person{firstName: "Tri", lastName: "Tran", contactInfo: contactInfo{phoneNumber: "395002471", countryCode: "+84"}}
	var thao person
	thao.firstName = "Thao"
	thao.lastName = "Nguyen"

	tri.updateFirstName("Frey")
	tri.printOutInfo()
	thao.printOutInfo()
}

func (p person) printOutInfo() {
	fmt.Println(p)
}

func (triPointer *person) updateFirstName(firstName string) {
	fmt.Println((*triPointer).firstName)
	(*triPointer).firstName = firstName
}
