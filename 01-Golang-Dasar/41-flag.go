package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port", 3306, "Put your database port")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
	fmt.Println("Port:", *port)
}
