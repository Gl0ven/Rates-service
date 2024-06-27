package handlers

import (
	grpc "Gl0ven/kata_projects/rates/internal/grpc/gen"
	"Gl0ven/kata_projects/rates/internal/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RatesHandler struct {
	service service.Service
	grpc.UnimplementedRatesServiceServer
}

func NewHandler(srv service.Service) grpc.RatesServiceServer {
	return &RatesHandler{
		service: srv,
		UnimplementedRatesServiceServer: grpc.UnimplementedRatesServiceServer{},
	}
}

func(h *RatesHandler) GetRates(ctx context.Context, emp *emptypb.Empty) (*grpc.RatesResponse, error) {
	rates, err := h.service.GetRates(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &grpc.RatesResponse{
		Timestamp: uint32(rates.Timestamp), 
		AskPrice: float32(rates.AskPrice), 
		BidPrice: float32(rates.BidPrice),
	}, nil
}