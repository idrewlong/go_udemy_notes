package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


func (eb englishBot) getGreeting() string {
	// custom logic for generating english greeting
	return "Hi There"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// If value isn't being used (eb englishBot), then we can remove the eb