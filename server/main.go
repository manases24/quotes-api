package main

import (
	"github.com/mnsh5/quotes/database"
	"github.com/mnsh5/quotes/routes"
)

func main() {
	database.ConnectDB()
	routes.Run()
}
