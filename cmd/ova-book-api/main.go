package main

import (
	"fmt"
	"log"
	"net"
	"os"

	cfg "github.com/ozonva/ova-book-api/internals/config"
	"github.com/ozonva/ova-book-api/internals/service"
	"github.com/ozonva/ova-book-api/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Не передан путь до файла конфигурации.")
		return
	}

	configPath := os.Args[1]
	config := cfg.ReadConfig(configPath)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("Ошибка при попытке занятия порта :%d, %v", config.Port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterBookServiceServer(grpcServer, service.NewBookApi())

	log.Printf("Запущен сервер на порту %d\n", config.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при старте grpc сервера на порту :%d: %v", config.Port, err)
	}

}
