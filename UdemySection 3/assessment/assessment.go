package main

import "fmt"

func main() {
	arrayOfInts := []int{1,2,3,4,5,6,7,8,9,10};
	for _, int := range arrayOfInts {
		if (int % 2 == 0) {
			fmt.Println("Even", int);
		} else {
			fmt.Println("Odd", int);
		}
	}
}