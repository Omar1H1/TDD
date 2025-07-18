package main

import "fmt"


// the string is the domain
func Hello(name string) string {
	return "Hello, " + name
}


/* 
   spreate domain from "side effects"
	 fmt.Println is a side effect "printing to stdout"
*/
func main() {
	fmt.Println(Hello("world"))
}
