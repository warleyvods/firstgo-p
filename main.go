package main

import (
	"firstgo-p/database"
	"firstgo-p/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
