package main

import "fmt"

func Ups() interface{} {
	return 1
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
