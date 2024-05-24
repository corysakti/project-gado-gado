package main

import "fmt"

func main() {
	// 1. string tidak wajib
	var name string

	name = "Eko Kurniawan Khannedy"
	fmt.Println(name)

	name = "EKo Khannedy"
	fmt.Println(name)

	// 2. string tidak wajib, tapi perlu di initialize
	var name2 = "test"

	name2 = "Eko Kurniawan Khannedy"
	fmt.Println(name2)

	name2 = "EKo Khannedy"
	fmt.Println(name2)

	// 3. shorthand tpai finalize ?
	name3 := "Test"
	fmt.Println(name3)

	name3 = "Eko Kurniawan Khannedy"
	fmt.Println(name3)

	name3 = "EKo Khannedy"
	fmt.Println(name3)

	// 4. var mirip struct ?
	var (
		firstName  = "Eko"
		middleName = "Kurniawan"
		lastName   = "khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
