package main

import (
	"time"

	client "github.com/KatsuyaAkasaka/grpc-gateway-sample/client"
	gateway "github.com/KatsuyaAkasaka/grpc-gateway-sample/gateway"
	server "github.com/KatsuyaAkasaka/grpc-gateway-sample/server"
)

func main() {
	serverPort := "19003"
	gwPort := "50000"
	go func() {
		server.Start(serverPort)
	}()
	time.Sleep(time.Second * 2)
	client.Call(serverPort)
	gateway.Start(serverPort, gwPort)
	return
}
