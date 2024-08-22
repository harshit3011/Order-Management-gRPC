package server

import (
	"log"
	"net/http"

	handler "github.com/harshit3011/Order-Management-gRPC/services/orders/handler/orders"
	"github.com/harshit3011/Order-Management-gRPC/services/orders/service"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (s *HttpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}