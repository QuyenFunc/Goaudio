package main

import (
	"database/sql"
	"log"

	"github.com/QuyenFunc/Goaudio/api"
	db "github.com/QuyenFunc/Goaudio/db/sqlc"
	"github.com/QuyenFunc/Goaudio/db/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot oad config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot connect sever:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot conecct server:", err)
	}
}
