package migrations

import (
	"context"
	"database/sql"
)

func UpSick(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		create table if not exists sick (
		    id integer primary key,
		    region_id integer,
		    name varchar(100),
		    count integer,
		    type_id integer,
		    registered_date timestamp,
		    created_at timestamp default current_timestamp,
			foreign key (region_id) references regions(id)
		);
	`)

	if err != nil {
		return err
	}
	return nil
}

func DownSick(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		drop table if exists sick;
	`)
	if err != nil {
		return err
	}
	return nil
}
