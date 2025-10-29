package main

import (
	"task4/config"
	"task4/internal/route"
)

func main() {
	config.InitDataBase()
	route.SetupRoutes()
}
