package standard_library

import (
	"fmt"
	"time"
)

func MainTime() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		fmt.Println(parse)
	} else {
		fmt.Println("Error", err.Error())
	}
}
