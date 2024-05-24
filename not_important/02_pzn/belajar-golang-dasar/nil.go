package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// return tipe data tidak bisa
// func Contoh(name string) string {
//	if name == "" {
//		return nil
//	} else {
//		return ""
//	}
//}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data Masih Kosong")
	} else {
		fmt.Println(data["name"])
	}

	fmt.Println(data["name"])
}
