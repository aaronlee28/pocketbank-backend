package main

import "C"
import (
	"fmt"
)

func main() {

	err := db.Connect()

	if err != nil {
		fmt.Println("failed to connect to db")

	}
	server.Init()
}
