package standard_library

import "fmt"

func MainFmt() {
	firstName := "Eko"
	lastName := "Khannedy"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
