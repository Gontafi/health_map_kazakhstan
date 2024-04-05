package db

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"real_time_health_map/internal/migrations"
)

func ConnectSqlLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:test.db?parseTime=true")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func UpMigrations(ctx context.Context, db *sql.DB) error {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = migrations.UpRegion(ctx, tx)
	if err != nil {
		return err
	}

	err = migrations.UpSick(ctx, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
