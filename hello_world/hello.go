package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// the string is the domain
func Hello(name, language string) string {

	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}


/* 
   spreate domain from "side effects"
	 fmt.Println is a side effect "printing to stdout"
*/
func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
