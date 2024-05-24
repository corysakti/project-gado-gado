package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	length := len(name)
	switch length > 5 {
	case true:
		fmt.Println("Nama kepanjangan")
	case false:
		fmt.Println("Nama benar")
	}

	switch {
	case length > 10:
		fmt.Println("Something went wrong!")
	case length > 5:
		fmt.Println("Nama kepanjangan")
	default:
		fmt.Println("Nama sudah benar")

	}
}
