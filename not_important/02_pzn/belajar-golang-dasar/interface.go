package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHelloW(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Eko"}
	sayHelloW(person)
}
