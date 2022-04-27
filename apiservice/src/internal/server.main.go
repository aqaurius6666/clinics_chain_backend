package main

import (
	"context"
	"fmt"

	"github.com/minh1611/go_structure/apiservice/src/internal/db"
	"github.com/minh1611/go_structure/apiservice/src/internal/db/my"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainServer, err := InitMainServer(ctx, ServerOptions{
		DBDsn: db.DBDsn(string("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	dbTables := mainServer.MainRepo.(*my.ServerCDBRepo).Interfaces

	mainServer.MainRepo.(*my.ServerCDBRepo).Db.AutoMigrate(dbTables...)
	mainServer.ApiServer.RegisterEndPoint()
	mainServer.ApiServer.G.Run(":1611")
}
