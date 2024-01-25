package database

import "fmt"

var connection string

func init() {
	fmt.Println("Calling database package...")
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}
