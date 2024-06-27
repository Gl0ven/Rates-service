package storage

import (
	"Gl0ven/kata_projects/rates/internal/models"
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Test_ratesStorage_SaveRates(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	mockSqlxDb := sqlx.NewDb(mockDb,"sqlmock")
	ctx := context.Background()
	rates1 := models.Rates{Timestamp: 1000000000, AskPrice:  90.1, BidPrice: 90.2}
	rates2 := models.Rates{}
	rs := &ratesStorage{db: mockSqlxDb}

	mock.ExpectExec("INSERT INTO rates").
	WithArgs(rates1.Timestamp, rates1.AskPrice, rates1.BidPrice).
	WillReturnResult(sqlmock.NewResult(1, 1)).
	WillReturnError(nil)

	mock.ExpectExec("INSERT INTO rates").
	WithArgs(rates2.Timestamp, rates2.AskPrice, rates2.BidPrice).
	WillReturnResult(sqlmock.NewResult(2, 1)).
	WillReturnError(fmt.Errorf("very bad error"))

	type args struct {
		ctx   context.Context
		rates models.Rates
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: ctx, rates: rates1},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{ctx: ctx, rates: rates2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rs.SaveRates(tt.args.ctx, tt.args.rates); (err != nil) != tt.wantErr {
				t.Errorf("ratesStorage.SaveRates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
