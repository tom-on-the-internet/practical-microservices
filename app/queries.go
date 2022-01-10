package app

import (
	"context"
	"fmt"
)

func (db *db) queryHomeViewData() (int, error) {
	var count int

	query := "SELECT COALESCE(SUM(view_count), 0) AS view_count FROM videos;"

	err := db.pool.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed queryHomeViewData: %w", err)
	}

	return count, nil
}
