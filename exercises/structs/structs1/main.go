// structs1
// Make me compile!
//
// I AM NOT DONE
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var person = Person{
		name: "John",
		age:  30,
	}

	fmt.Printf("Person %s and age %d", person.name, person.age)
}
