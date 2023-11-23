package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	hostname, err := os.Hostname()

	fmt.Println(args)

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err.Error())
	}

}
