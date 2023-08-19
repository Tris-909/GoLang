package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "Hi There"
}

func main() {
	eb := engBot{}

	printGreeting(eb)
}
