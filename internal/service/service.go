package service

import (
	"context"
	"database/sql"
	"real_time_health_map/internal/models"
	"real_time_health_map/internal/repo"
)

func GetStats(ctx context.Context, db *sql.DB, options *repo.SickOptions) (map[int16]models.Statistics, error) {
	stats, err := repo.GetStatsFromSickTable(ctx, db, options)
	if err != nil {
		return nil, err
	}
	// future logic

	return stats, err
}

func InsertStat(ctx context.Context, db *sql.DB, sick models.Sick) error {
	err := repo.InsertStatistics(ctx, db, sick)
	if err != nil {
		return err
	}

	return nil
}
