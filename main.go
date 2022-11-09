package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/wizlif/simplebank/api"
	db "github.com/wizlif/simplebank/db/sqlc"
	"github.com/wizlif/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("could not start server: ", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server", config.ServerAddress)
	}

}
