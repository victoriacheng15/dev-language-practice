package main

import "fmt"

// Day 40: Hello World Test
// https://github.com/quii/learn-go-with-tests/blob/main/hello-world.md

// const Spanish = "Spanish"
// const French = "French"
// const Mandarin = "Mandarin"
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "
// const mandarinHelloPrefix = "Ni Hao, "

// can write constants in a single block or each line like above
const (
	Spanish             = "Spanish"
	French              = "French"
	Mandarin            = "Mandarin"
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	mandarinHelloPrefix = "Ni Hao, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch language {
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	case Mandarin:
		prefix = mandarinHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Victoria", Mandarin))
}
