package standard_library

import (
	"fmt"
	"time"
)

func MainDuration() {
	var duration1 = 100 * time.Second
	var duration2 = 10 * time.Millisecond
	var duration3 = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration : %d\n", duration3)

}
