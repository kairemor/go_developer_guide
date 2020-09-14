package main

import "fmt"

type contactInfo struct {
	email   string
	zipInfo int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	mor := person{
		firstName: "Mor",
		lastName:  "Kaire",
		contactInfo: contactInfo{
			email:   "kairemor@gmail.com",
			zipInfo: 54666,
		},
	}
	// In go when we use a receiver with pointer of some type
	// it can also use with the receiver of only this type
	// but not a pointer
	// morPointer := &mor
	// morPointer.updateName("Lamine")

	mor.updateName("Cheikh")

	mor.print()

	// go will automatically allow a "zero" value when we declare variable without initial value
	var lamine person
	lamine.firstName = "Lamine"

	// A slice is working in a very different way
	// when we pass a slice as a parameter of a function
	// it will be passed by reference contrary to what go
	// is doing normally thus slice are reference type contrary to struct

	/*
	*Otherwise slice are also in reality  pass by value but when we declare a slice
	*go declare two things :
	*	an array of the value of the slice
	*	an variable which reference the first element of the array, the length of the array and the capacity
	 */
	greeting := []string{"Hi", "Hello", "how", "nice"}

	update(greeting)
	fmt.Println(greeting)
}

func (p person) print() {
	// below allow to print key and value at the same time
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func update(s []string) {
	s[0] = "changed"
}

/*
in go there are reference type and value
to modify the value type in a function we need to use pointer
and for reference type we don't need it
*/
