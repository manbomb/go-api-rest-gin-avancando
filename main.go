package main

import (
	"github.com/manbomb/api-go-gin/database"
	"github.com/manbomb/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
