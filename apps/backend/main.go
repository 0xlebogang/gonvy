package main

import (
	"log"

	"github.com/0xlebogang/gonvy/backend/internal/config"
	"github.com/0xlebogang/gonvy/backend/internal/database"
	"github.com/0xlebogang/gonvy/backend/internal/server"
)

func main() {
	conf := config.Load()

	db, err := database.ConnectPg(conf.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Fatal(err.Error())
		}
	}()

	svr := server.New(conf, db)
	if err := svr.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
