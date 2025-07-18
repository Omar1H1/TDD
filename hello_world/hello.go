package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


// the string is the domain
func Hello(name, language string) string {

	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}


func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}


/* 
spreate domain from "side effects"
fmt.Println is a side effect "printing to stdout"
*/
func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
