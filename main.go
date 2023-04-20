package main

import (
	"database/sql"
	"github.com/avosaga/go-simplebank/api"
	db "github.com/avosaga/go-simplebank/db/sqlc"
	"github.com/avosaga/go-simplebank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("couldn't read config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("couldn't connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("couldn't create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("couldn't start server:", err)
	}
}
