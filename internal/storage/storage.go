package storage

import (
	"context"
	"database/sql"
	"fmt"

	"mail-sender/internal/config"
)

type Storage struct {
	config *config.DatabaseConfig
}

func New(dbConfig *config.DatabaseConfig) *Storage {
	return &Storage{
		config: dbConfig,
	}
}

const connectionStringPattern = "postgres://%s:%s@%s:%d/%s?sslmode=disable"

func buildConnectionString(dbConfig *config.DatabaseConfig) string {
	return fmt.Sprintf(connectionStringPattern, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
}

func connect(dbConfig *config.DatabaseConfig) (*sql.DB, error) {
	connection, err := sql.Open("postgres", buildConnectionString(dbConfig))
	if err != nil {
		return nil, fmt.Errorf("connect to database failed: %w", err)
	}

	if err = connection.Ping(); err != nil {
		return nil, fmt.Errorf("connect to database failed: %w", err)
	}

	return connection, nil
}

const insertEventQuery = "INSERT INTO public.events (data, created_at) VALUES ($1, current_timestamp)"

func (i *Storage) InsertEvent(ctx context.Context, data []byte) error {
	conn, err := connect(i.config)
	if err != nil {
		return fmt.Errorf("insert event failed: %w", err)
	}
	defer conn.Close()

	result, err := conn.ExecContext(ctx, insertEventQuery, data)
	if err != nil {
		return fmt.Errorf("insert event failed: %w", err)
	}

	return nil
}

const selectEventQuery = "SELECT data FROM public.events LIMIT $1"

func (i *Storage) SelectEvent(ctx context.Context, batchSize int) ([][]byte, error) {
	conn, err := connect(i.config)
	if err != nil {
		return nil, fmt.Errorf("select event failed: %w", err)
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, selectEventQuery, batchSize)
	if err != nil {
		return nil, fmt.Errorf("select event failed: %w", err)
	}

	var result [][]byte
	for rows.Next() {
		var tmpData []byte

		rows.Scan(&tmpData)

		result = append(result, tmpData)
	}

	return result, nil
}

const removeEventQuery = "DELETE FROM public.events WHERE id = $1"

func (i *Storage) RemoveEvent(ctx context.Context, id int) error {
	conn, err := connect(i.config)
	if err != nil {
		return fmt.Errorf("remove event failed: %w", err)
	}
	defer conn.Close()

	result, err := conn.ExecContext(ctx, removeEventQuery, id)
	if err != nil {
		return fmt.Errorf("remove event failed: %w", err)
	}

	return nil
}
