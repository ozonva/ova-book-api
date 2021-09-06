package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	cfg "github.com/ozonva/ova-book-api/internals/config"
	"github.com/ozonva/ova-book-api/internals/db"
	producer "github.com/ozonva/ova-book-api/internals/kafka"
	"github.com/ozonva/ova-book-api/internals/metrics"
	"github.com/ozonva/ova-book-api/internals/service"
	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Не передан путь до файла конфигурации.")
		return
	}

	configPath := os.Args[1]
	appConfig := cfg.ReadConfig(configPath)

	connection := db.Connect(
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			appConfig.DBUser, appConfig.DBPwd, appConfig.DBHost, appConfig.DBName,
		))
	repo := db.CreateRepo(connection)

	kafkaProducer := producer.New(appConfig.KafkaServers)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.Port))
	if err != nil {
		log.Fatalf("Ошибка при попытке занятия порта :%d, %v", appConfig.Port, err)
	}

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"ova-book-api",
		config.Logger(jaeger.StdLogger),
	)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	createdCounter := metrics.New("createdCounter", "Счетчик успешно созданных записей")
	updatededCounter := metrics.New("updatededCounter", "Счетчик успешно обновленных записей")
	deletedCounter := metrics.New("deletedCounter", "Счетчик успешно удаленных записей")

	grpcServer := grpc.NewServer()
	api.RegisterBookServiceServer(
		grpcServer,
		service.NewBookApi(repo, &kafkaProducer, &createdCounter, &updatededCounter, &deletedCounter),
	)

	go http.ListenAndServe(
		fmt.Sprintf(appConfig.PrometheusExporter),
		promhttp.Handler(),
	)

	log.Printf("Запущен сервер на порту %d\n", appConfig.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при старте grpc сервера на порту :%d: %v", appConfig.Port, err)
	}

}
