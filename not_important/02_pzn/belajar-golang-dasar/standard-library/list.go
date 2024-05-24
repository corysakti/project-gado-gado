package standard_library

import (
	"container/list"
	"fmt"
)

func MainList() {
	var data = list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	var head = data.Front()
	fmt.Println(head.Value) // eko

	next := head.Next()
	fmt.Println(next.Value) // kurniawan

	next = head.Next() // khannedy
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
