package repo

import (
	"context"
	"database/sql"
	"real_time_health_map/internal/models"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type SickOptions struct {
	DateFrom int64
	SickType string
}

func GetStatsFromSickTable(ctx context.Context, db *sql.DB, opts *SickOptions) (map[int16]models.Statistics, error) {
	statistics := make(map[int16]models.Statistics)
	if opts == nil {
		opts = new(SickOptions)
		opts.DateFrom = time.Now().Add(-1 * time.Hour * 24 * 30).Unix()
	}

	query := sq.Select("region_id", "type_id", "name", "count").From("sick").
		Where(sq.GtOrEq{"registered_date": opts.DateFrom})

	if opts.SickType != "" {
		query = query.Where(sq.Eq{"name": opts.SickType})
	}

	stmt, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, nil
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count    int
			typeID   int
			regionID int16
			name     string
		)

		err := rows.Scan(&regionID, &typeID, &name, &count)
		if err != nil {
			return nil, err
		}
		stats, ok := statistics[regionID]
		if !ok {
			statistics[regionID] = models.Statistics{}
		} else if opts.SickType != "" {
			stats.SickName = opts.SickType
		} else {
			stats.SickName = "All"
		}

		switch typeID {
		case 3:
			stats.Dead += count
		case 2:
			stats.Cured += count
		case 1:
			stats.Sick += count
		}

		statistics[regionID] = stats
	}

	return statistics, nil
}

func InsertStatistics(ctx context.Context, db *sql.DB, stat models.Sick) error {

	query := sq.Insert("sick").
		Columns("region_id", "sick_name", "count", "type_id").
		Values(stat.RegionID, stat.SickName, stat.Count, stat.TypeID)

	stmt, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return nil
}
