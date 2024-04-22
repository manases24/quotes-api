package main

import (
	"github.com/mnsh5/quotes/database"
	"github.com/mnsh5/quotes/routes"
)

type User struct {
	Name    string
	Country string
}

func main() {
	database.ConnectDB()
	routes.Run()
}
