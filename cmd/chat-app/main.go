package main

import (
	"fmt"
	"log"

	"github.com/ShariUllas/chat-app/pkg/config"
	"github.com/ShariUllas/chat-app/pkg/handlers"
	"github.com/ShariUllas/chat-app/pkg/services"
)

func main() {
	appConfig := config.New()
	dbConn := appConfig.GetDBConnectionString()
	db, err := appConfig.GetDBConnection(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	service := services.New(appConfig)
	engine := handlers.New(appConfig, service)
	engine.Run()
}
