package standard_library

import (
	"fmt"
	"slices"
)

func MainSlice() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Index(names, "Eko"))
	fmt.Println(slices.Index(names, "Paul"))
}
