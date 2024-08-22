package main

import (
    "log"
	"github.com/harshit3011/Order-Management-gRPC/services/orders/server"
)

func main() {
    httpServer := server.NewHttpServer(":8000")
    go func() {
        if err := httpServer.Run(); err != nil {
            log.Fatalf("HTTP server failed: %v", err)
        }
    }()

    grpcServer := server.NewGRPCServer(":9000")
    if err := grpcServer.Run(); err != nil {
        log.Fatalf("gRPC server failed: %v", err)
    }
}
