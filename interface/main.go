package main

import "fmt"

// interface make easy to re use code 
type englishBot struct{}
type spanishBot struct{}
// interface declaration 
// this say to every other type we have a custom type called bot/*

/*
 * if there are many function inside interface a type should also have all this function to be interface type
 *there are concrete and interface type . we cannot declare variable as type interface
	* interfave are nor go don't support generic type like in java 
	* interface are implicit type
	* we don't have to explicitly say that a concrete type is a honorary of a interface type but go handle it it's implicit 
	* interfave are a contract to help us manage type 
  * interface are tough step1 is understanding how to read them in the documentation official
*/
type bot interface {
	//if a type in the program have a getGreeting function returning a string
	// this type us now also of type our custom type.bot they are honorary member 
	//of this type
	getGreeting() string
}

func main() {
	eb := englishBot{}
	es := spanishBot{}

	printGreeting(eb)
	printGreeting(es)
}

//
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//if we don't use the receiver variable we can delete it like in the two follow function
func (englishBot) getGreeting() string {
	//custom logic tongenerate english greering 
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	//the same also custom logic to generate  spanish greeting
	return "Hola!"
}
