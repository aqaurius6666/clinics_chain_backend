package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/psql"
	"github.com/minh1611/clinics_chain_management/apiservice/src/services/authservice"
)

// change to  main
func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load("../../deploy/dev/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Start http server
	httpListener, err := net.Listen("tcp", ":"+os.Getenv("HTTP_PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	mainServer, err := InitMainServer(ctx, ServerOptions{
		DBDsn:           db.DBDsn(os.Getenv("CONFIG_DB_URI")),
		AuthServiceAddr: authservice.AuthServiceAddr(os.Getenv("CONFIG_AUTHSERVICE_ADDRESS")),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Auto migrate entity to sql "Table"
	dbTables := mainServer.MainRepo.(*psql.ServerCDBRepo).Interfaces
	// doc for dbTable...: https://go.dev/ref/spec#Passing_arguments_to_..._parameters
	err = mainServer.MainRepo.(*psql.ServerCDBRepo).Db.AutoMigrate(dbTables...)
	if err != nil {
		fmt.Println(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Create route
		mainServer.ApiServer.RegisterEndPoint()

		if err := http.Serve(httpListener, mainServer.ApiServer.G); err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
		fmt.Println("App is running at port " + os.Getenv("HTTP_PORT"))
	}()
	wg.Wait()
}
