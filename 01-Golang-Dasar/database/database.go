package database

import "fmt"

var connection string

func init() {
	fmt.Println("Run Init")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
