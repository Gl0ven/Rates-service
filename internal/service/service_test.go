package service

import (
	"Gl0ven/kata_projects/rates/internal/models"
	pmocks "Gl0ven/kata_projects/rates/internal/provider/garantex/mocks"
	smocks "Gl0ven/kata_projects/rates/internal/storage/mocks"
	"context"
	"fmt"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func Test_ratesService_GetRates(t *testing.T) {
	mockStorage := smocks.NewStorage(t)
	mockProvider := pmocks.NewGarantexApi(t)
	testLogger := zap.NewExample()
	rates1 := models.Rates{Timestamp: 1000000000, AskPrice:  90.1, BidPrice: 90.2}
	rates2 := models.Rates{Timestamp: 10000000001, AskPrice:  90.2, BidPrice: 90.1}
	rates3 := models.Rates{}
	err := fmt.Errorf("very bad error")
	ctx := context.Background()

	mockProvider.On("GetRates").Return(rates1, nil).Times(1)
	mockProvider.On("GetRates").Return(rates2, nil).Times(1)
	mockProvider.On("GetRates").Return(rates3, err)
	mockStorage.On("SaveRates", ctx, rates1).Return(nil)
	mockStorage.On("SaveRates", ctx, rates2).Return(err)
	
	rs := &ratesService{
		storage:  mockStorage,
		provider: mockProvider,
		logger:   testLogger,
	}

	tests := []struct {
		name    string
		args    context.Context
		want    models.Rates
		wantErr bool
	}{
		{
			name: "success",
			args: ctx, 
			want: rates1, 
			wantErr: false,
		},
		{
			name: "fail1", 
			args: ctx,
			want: rates3, 
			wantErr: true,
		},
		{
			name: "fail2", 
			args: ctx,
			want: rates3, 
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rs.GetRates(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ratesService.GetRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ratesService.GetRates() = %v, want %v", got, tt.want)
			}
		})
	}
}
