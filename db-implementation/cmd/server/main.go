package main

import (
	"db-implementation/cmd/server/routes"
	"db-implementation/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
