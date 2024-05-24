package standard_library

import (
	"flag"
	"fmt"
)

func MainFlag() {
	var username = flag.String("username", "root", "database username")
	var password = flag.String("password", "root", "database password")
	var host = flag.String("host", "root", "database host")
	var port = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}
