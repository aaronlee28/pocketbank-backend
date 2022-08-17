package main

import "C"
import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/aaron-lee/final-project-backend/server"
)

func main() {

	err := db.Connect()

	if err != nil {
		fmt.Println("failed to connect to db")

	}

	server.Init()
}
