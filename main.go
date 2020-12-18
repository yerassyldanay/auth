package main

import (
	"auth/api"
	"auth/utils/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	// loads env variables from the directory of ./environment
	configuration, err := config.LoadConfig("./environment")
	if err != nil {
		log.Fatal("could not load env variables: ", err)
	}

	// connect to database
	connection, err := sql.Open(configuration.DbName, configuration.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer connection.Close()

	// create a new server
	server := api.NewServer(connection)

	// start a server
	err = server.Start("0.0.0.0:8100")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
