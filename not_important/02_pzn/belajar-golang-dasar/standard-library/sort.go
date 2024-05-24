package standard_library

import "sort"

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (receive UserSlice) Len() int {
	return len(receive)
}

func (receive UserSlice) Less(i, j int) bool {
	return receive[i].Age < receive[i].Age
}

func (receive UserSlice) Swap(i, j int) {
	receive[i], receive[j] = receive[j], receive[i]
}

func MainSort() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Anwar", 20},
	}
	sort.Sort(UserSlice(users))
}
