package service

import (
	"context"
	"test_lamoda/internal/db"
	proto "test_lamoda/internal/service/proto/gen"
)

func (s *Server) ReserveProducts(ctx context.Context, req *proto.ReserveProductsRequest) (*proto.ReserveProductsResponse, error) {
	err := db.ReserveProduct(s.db, req.UniqueCodes)
	if err != nil {
		return nil, err
	}

	return &proto.ReserveProductsResponse{
		Message: "Products reserved successfully",
	}, nil
}

func (s *Server) ReleaseReservations(ctx context.Context, req *proto.ReleaseReservationsRequest) (*proto.ReleaseReservationsResponse, error) {
	err := db.ReleaseReservation(s.db, req.UniqueCodes)
	if err != nil {
		return nil, err
	}

	return &proto.ReleaseReservationsResponse{
		Message: "Reservations released successfully",
	}, nil
}

func (s *Server) GetRemainingProducts(ctx context.Context, req *proto.GetRemainingProductsRequest) (*proto.GetRemainingProductsResponse, error) {
	remainingCount, err := db.GetRemainingProducts(s.db, req.GetStoreId())
	if err != nil {
		return nil, err
	}

	return &proto.GetRemainingProductsResponse{
		Count: int32(remainingCount),
	}, nil
}
