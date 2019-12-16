package main

import "fmt"

type greetable interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	greet(eb)
	greet(sb)
}

func greet(g greetable) {
	fmt.Println(g.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hey there!"
}

func (spanishBot) getGreeting() string {
	return "Hola! Que tal?"
}
