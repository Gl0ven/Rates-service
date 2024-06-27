package handlers

import (
	grpc "Gl0ven/kata_projects/rates/internal/grpc/gen"
	"Gl0ven/kata_projects/rates/internal/models"
	"Gl0ven/kata_projects/rates/internal/service/mocks"
	"context"
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRatesHandler_GetRates(t *testing.T) {
	mockService := mocks.NewService(t)
	err := fmt.Errorf("very bad error")
	ctx := context.Background()
	emp := &emptypb.Empty{}
	rh := RatesHandler{
		service: mockService,
		UnimplementedRatesServiceServer: grpc.UnimplementedRatesServiceServer{},
	}
	rates1 := models.Rates{Timestamp: 1000000000, AskPrice:  90.1, BidPrice: 90.2}
	rates2 := models.Rates{}

	mockService.On("GetRates", ctx).Return(rates1, nil).Times(1)
	mockService.On("GetRates", ctx).Return(rates2, err)
	
	type args struct {
		ctx context.Context
		empty *emptypb.Empty
	}

	tests := []struct {
		name    string
		args args
		want    *grpc.RatesResponse
		wantErr bool
	}{
		{
			name: "success", 
			args: args{ctx: ctx, empty: emp}, 
			want: &grpc.RatesResponse{Timestamp: 1000000000, AskPrice: 90.1, BidPrice: 90.2}, 
			wantErr: false,
		},
		{
			name: "fail", 
			args: args{ctx: ctx, empty: emp}, 
			want: nil, 
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rh.GetRates(tt.args.ctx, tt.args.empty)
			if (err != nil) != tt.wantErr {
				t.Errorf("RatesHandler.GetRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RatesHandler.GetRates() = %v, want %v", got, tt.want)
			}
		})
	}
}
