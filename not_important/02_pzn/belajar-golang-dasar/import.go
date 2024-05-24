package main

import (
	"belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println("Eko")
	helper.SayHello("Test")
	fmt.Println(database.GetDatabase())
}
