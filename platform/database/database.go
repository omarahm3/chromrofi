package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

type Database struct {
	databaseUrl string
	Dbx         *sqlx.DB
}

func NewDatabase(databaseUrl string) *Database {
	return &Database{
		databaseUrl: databaseUrl,
	}
}

func (d *Database) Connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, d.databaseUrl)
	if err != nil {
		return err
	}

	dbx.DB.SetMaxOpenConns(300)
	dbx.DB.SetMaxIdleConns(50)
	dbx.DB.SetConnMaxLifetime(0)

	d.Dbx = dbx
	return nil
}

func (d *Database) Close() error {
	return d.Dbx.Close()
}
