// Create a function using interface and struct to calculate the area of a rectangle

package main

import "fmt"

type Area interface {
	calculateArea() int
}

type rectangle struct {
	length, width int
}

func calculate(a Area) {
	fmt.Println(a.calculateArea())
}

func (r rectangle) calculateArea() int {
	return r.length * r.width
}

func main() {
	stadium := rectangle{length: 10, width: 8}
	calculate(stadium)
}
