package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/minh1611/go_structure/apiservice/src/internal/db"
	"github.com/minh1611/go_structure/apiservice/src/internal/db/my"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get variable from .env
	err := godotenv.Load("../../deploy/dev/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	mainServer, err := InitMainServer(ctx, ServerOptions{
		DBDsn: db.DBDsn(os.Getenv("CONFIG_DB_URI")),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Auto migrate entity to sql "Table"
	dbTables := mainServer.MainRepo.(*my.ServerCDBRepo).Interfaces
	// doc for dbTable...: https://go.dev/ref/spec#Passing_arguments_to_..._parameters
	mainServer.MainRepo.(*my.ServerCDBRepo).Db.AutoMigrate(dbTables...)

	// Create route
	mainServer.ApiServer.RegisterEndPoint()

	// Run server
	mainServer.ApiServer.G.Run(":" + os.Getenv("PORT"))
}
