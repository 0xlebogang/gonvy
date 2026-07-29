package main

import (
	"log"

	"github.com/0xlebogang/gonvy/api/internal/config"
	"github.com/0xlebogang/gonvy/api/internal/server"
	"github.com/0xlebogang/gonvy/api/internal/storage/database"
)

func main() {
	conf := config.Load()

	dbStore := database.New(&conf.DatabaseConf)
	dbConn, err := dbStore.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		if err := dbStore.Close(); err != nil {
			log.Fatalf("Database disconnection failed: %v", err)
		}
	}()

	dbStore.RunMigrations()

	server := server.New(&conf.ServerConf, dbConn)
	if err := server.Start(); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
