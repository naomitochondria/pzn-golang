package main

import (
	"fmt"
	"go-dasar/database"
	_ "go-dasar/helper"
	"testing"
)

func TestPackageInitialization(t *testing.T) {
	// Import and package initialization
	connection := database.GetConnection()
	fmt.Println("Connection: ", connection)

	// Ignore but init package
	// fmt.Println(helper.Multiply())
}
