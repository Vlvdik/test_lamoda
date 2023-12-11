package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	proto "test_lamoda/internal/service/proto/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	db *sql.DB
	proto.UnimplementedProductServiceServer
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) RunGateway(grpcServerAddr, gatewayAddr string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterProductServiceHandlerFromEndpoint(ctx, mux, grpcServerAddr, opts)
	if err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("gRPC gateway is listening on addr %s and forwarding requests to gRPC server at %s\n", gatewayAddr, grpcServerAddr)
	if err := http.ListenAndServe(gatewayAddr, mux); err != nil {
		log.Fatalf("failed to serve HTTP gateway: %v", err)
	}
}
