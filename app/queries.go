package app

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
)

type versionConflictError struct {
	expectedVersion int
	streamVersion   int
}

func (v versionConflictError) Error() string {
	return fmt.Sprintf("Version Error: expected %d but got %d.", v.expectedVersion, v.streamVersion)
}

func (db *db) queryHomeViewData() (int, error) {
	var count int

	query := "SELECT COALESCE(SUM(view_count), 0) AS view_count FROM videos;"

	err := db.pool.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed queryHomeViewData: %w", err)
	}

	return count, nil
}

func (msgStore *messageStore) write(streamName string, msg message, expectedVersion int) error {
	query := "SELECT message_store.write_message($1, $2, $3, $4, $5, $6)"

	tag, err := msgStore.pool.Exec(context.Background(), query, msg.id, streamName, msg.msgType, msg.data, msg.metadata, expectedVersion)
	if err != nil {
		r := regexp.MustCompile(`\Stream Version: (.*?)\)`)
		matches := r.FindStringSubmatch(err.Error())

		isStreamVersionError := len(matches) > 1

		if !isStreamVersionError {
			return fmt.Errorf("failed message store write: %w", err)
		}

		streamVersion, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("failed message store write: %w", err)
		}

		return versionConflictError{
			expectedVersion: expectedVersion,
			streamVersion:   streamVersion,
		}
	}

	_ = tag

	return nil
}
