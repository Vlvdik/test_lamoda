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
	"time"

	"google.golang.org/grpc"
)

func main() {
	config := config.LoadConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", config.DBUser, config.DB, config.DBPassword, config.DBHost))
	for err != nil {
		log.Printf("Waiting for the database to be available...")
		time.Sleep(1 * time.Second)
		db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", config.DBUser, config.DB, config.DBPassword, config.DBHost))
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

	server.RunGateway(config.GRPCAddr, config.HTTPGatewayAddr)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	slog.Info("Shutting down gracefully")

	grpcServer.GracefulStop()
}
