package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainServer, err := InitMainServer(ctx)
	if err != nil {
		return
	}
	mainServer.ApiServer.RegisterEndPoint()
	mainServer.ApiServer.G.Run(":1611")
}
