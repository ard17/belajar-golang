package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("aki"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("ari"))
	fmt.Println(regex.MatchString("aTi"))

	fmt.Println(regex.FindAllString("ali aki aji ako ayi fii", -1))
}
