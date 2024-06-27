package db

import (
	"Gl0ven/kata_projects/rates/config"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(dbConf config.DB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
	 	dbConf.Host, dbConf.Port, dbConf.Name, dbConf.User, dbConf.Password,
	)
	dbRaw, err := sql.Open(dbConf.Driver, dsn)
	if err != nil {
		return nil, err
	}
	db := sqlx.NewDb(dbRaw, dbConf.Driver)

	return db, nil
}