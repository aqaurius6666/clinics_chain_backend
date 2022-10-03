package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/minh1611/go_structure/authservice/src/internal/db"
	"github.com/minh1611/go_structure/authservice/src/internal/db/my"
	"github.com/minh1611/go_structure/authservice/src/pb"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup

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
	// doc for dbTable...: https://go.dev/ref/spec#Passing_arguments_to_..._parameters
	dbTables := mainServer.MainRepo.(*my.ServerCDBRepo).Interfaces
	mainServer.MainRepo.(*my.ServerCDBRepo).Db.AutoMigrate(dbTables...)

	// Start GRPC Server
	grpcListener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = grpcListener.Close() }()

	var srv *grpc.Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.Serve(grpcListener); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		// Register
		pb.RegisterAuthServiceServer(srv, mainServer.ApiServer)
	}()

	wg.Wait()
}
