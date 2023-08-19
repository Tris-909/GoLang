// Exercise from myself

// Create an interface to decide if a human is an adult or not and print out their name as well
// STEPS :
// 1. Create a struct for human
// 2. Create isAdult function that validToCalculateIfTheyAreAnAdult need to satisfy
// 3. Define functions like isYearOfBirthExists and getName
// 4. instances of human and call isAdult
// 5. See if the results are correct

package main

import (
	"fmt"
)

type human struct {
	name        string
	height      int
	gender      string
	yearOfBirth int
}

type validToCalculateIfTheyAreAnAdult interface {
	isYearOfBirthExists() bool
	getName() string
}

func (h human) isYearOfBirthExists() bool {
	if h.yearOfBirth == 0 {
		return false
	}

	return true
}

func (h human) getName() string {
	return h.name
}

func isAdult(v validToCalculateIfTheyAreAnAdult) {
	YoBValid := v.isYearOfBirthExists()
	nameValid := v.getName()

	if nameValid == "" {
		return
	}

	if YoBValid {
		fmt.Println(v.getName(), "Is an adult")
	} else {
		fmt.Println(v.getName(), "NOT an adult")
	}
}

func main() {
	tri := human{name: "Tri", height: 173, gender: "male", yearOfBirth: 1999}
	randomDude := human{name: "RandomDude", height: 150, gender: "male"}
	dudeWithNoName := human{height: 170, gender: "gay", yearOfBirth: 2000}

	isAdult(tri)
	isAdult(randomDude)
	isAdult(dudeWithNoName)
}
