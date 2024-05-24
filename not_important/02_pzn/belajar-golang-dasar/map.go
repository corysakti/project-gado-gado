package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Eko"
	//person["address"] = "Subang"

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//make map
	book := make(map[string]string)
	book["title"] = "Buku Go-lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "Ups"

	delete(book, "wrong")
	fmt.Println(book)
}
