package standard_library

import (
	"encoding/base64"
	"fmt"
)

func MainBase64() {
	value := "Eko Kurniawan Khannedy"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
