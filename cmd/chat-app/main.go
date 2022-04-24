package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/ShariUllas/chat-app/pkg/config"
	"github.com/ShariUllas/chat-app/pkg/handlers"
	"github.com/ShariUllas/chat-app/pkg/services"
)

var migDir = "/tmp/migrations"

func main() {
	appConfig := config.New()
	dbConn := appConfig.GetDBConnectionString()
	db, err := appConfig.GetDBConnection(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	appConfig.MigrateDB(dbConn, migDir)
	fmt.Println(db)
	service := services.New(appConfig, db)
	engine := handlers.New(appConfig, service)
	ch := make(chan error)

	go func() { ch <- RunWebSocket(appConfig.WSPort) }()
	go func() { ch <- engine.Run(appConfig.Host, appConfig.APIPort) }()
	channel_err := <-ch
	log.Fatalln(errors.Wrap(channel_err, "service failed"))
}
