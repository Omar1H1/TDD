package main

import "fmt"

const englishHelloPrefix = "Hello, "

// the string is the domain
func Hello(name string) string {

	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name
}


/* 
   spreate domain from "side effects"
	 fmt.Println is a side effect "printing to stdout"
*/
func main() {
	fmt.Println(Hello("world"))
}
