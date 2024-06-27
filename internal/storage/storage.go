package storage

import (
	"Gl0ven/kata_projects/rates/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type Storage interface {
	SaveRates(ctx context.Context, rates models.Rates) error
}

type ratesStorage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Storage {
	return &ratesStorage{
		db: db,
	}
}

func(s *ratesStorage) SaveRates(ctx context.Context, rates models.Rates) error {
	_, err := s.db.Exec(
		"INSERT INTO rates (unix_timestamp, ask_price, bid_price) VALUES ($1, $2, $3)",
		rates.Timestamp, rates.AskPrice, rates.BidPrice,
	)
	return err
}