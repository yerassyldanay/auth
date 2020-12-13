package main

import (
	"auth/api"
	database "auth/model/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var (
	dbName = "postgres"
	dbSource = "postgres://simple:simple@localhost:8101/simple?sslmode=disable"
)

func main() {
	connection, err := sql.Open(dbName, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	/*
		TODO: load env values
	*/

	var query = database.New(connection)
	server := api.NewServer(query)

	err = server.Start("0.0.0.0:8100")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
