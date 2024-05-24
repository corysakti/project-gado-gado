package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko Kurniawan Khannedy"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko)

	//struct literal
	joko := Customer{
		Name:    "A",
		Address: "B",
		Age:     15,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	eko.sayHello("Agus")
	joko.sayHello("Agus")
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}
