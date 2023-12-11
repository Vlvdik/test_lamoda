package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"test_lamoda/config"
	"test_lamoda/internal/service"
	proto "test_lamoda/internal/service/proto/gen"

	"google.golang.org/grpc"
)

func main() {
	config := config.LoadConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", "", "", ""))
	if err != nil {
		slog.Error("failed to connect to the database: ", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	server := service.NewServer(db)

	proto.RegisterProductServiceServer(grpcServer, server)

	go func() {
		listener, err := net.Listen("tcp", config.GRPCAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
			return
		}

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	server.RunGateway(config.GRPCAddr, config.HTTPGatewayPort)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала о завершении работы
	<-stop
	slog.Info("Shutting down gracefully")

	grpcServer.GracefulStop()
}
