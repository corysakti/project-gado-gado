package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "eko" {
		return &notFoundError{"data nnot found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		switch finalError := err.(type) {
		case *validationError:

			fmt.Println("Validation Error : ", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error : ", finalError.Error())
		default:
			fmt.Println("unknown error : ", finalError.Error())
		}
	} else {
		fmt.Println("Success!")
	}
}
