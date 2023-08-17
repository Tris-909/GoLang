package main

import "fmt"

func main() {
	secColors := make(map[string]string)
	secColors["white"] = "#fff"

	for key, value := range secColors {
		fmt.Println(key, value)
	}

	delete(secColors, "white")
}
