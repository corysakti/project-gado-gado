package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return "Anjing" == name
	}
	registerUser("Eko", blackList)

	//atau langsung
	registerUser("Eko", func(name string) bool {
		return "Anjing" == name
	})
}
