package migrations

import (
	"context"
	"database/sql"
)

func UpRegion(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		create table if not exists region (
		    id integer primary key,
		    name varchar(100)
		);
	`)

	if err != nil {
		return err
	}
	return nil
}

func DownRegion(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		drop table if exists region;
	`)
	if err != nil {
		return err
	}
	return nil
}
