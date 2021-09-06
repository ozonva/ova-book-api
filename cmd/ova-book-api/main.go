package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	cfg "github.com/ozonva/ova-book-api/internals/config"
	"github.com/ozonva/ova-book-api/internals/db"
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

	connection := db.Connect(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.DBUser, config.DBPwd, config.DBHost, config.DBName))
	repo := db.CreateRepo(connection)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("Ошибка при попытке занятия порта :%d, %v", config.Port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterBookServiceServer(grpcServer, service.NewBookApi(repo))

	log.Printf("Запущен сервер на порту %d\n", config.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при старте grpc сервера на порту :%d: %v", config.Port, err)
	}

}
