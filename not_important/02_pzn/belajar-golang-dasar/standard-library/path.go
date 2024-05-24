package standard_library

import (
	"fmt"
	"path"
)

func MainPath() {
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
