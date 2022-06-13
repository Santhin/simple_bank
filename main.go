package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/santhin/simple_bank/api"
	db "github.com/santhin/simple_bank/db/sqlc"
	"github.com/santhin/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
