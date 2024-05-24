package main

import "fmt"

func main() {

	type NoKtp string

	var ktpEko NoKtp = "123"

	var contoh string = "321"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
}
