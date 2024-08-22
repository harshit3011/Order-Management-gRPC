package main

import "github.com/harshit3011/Order-Management-gRPC/services/kitchen/client"



func main() {
	httpServer := client.NewHttpServer(":1000")
	httpServer.Run()
}