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
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/mailservice/src/services/authservice"
)

// change to  main
func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get variable from .env
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
