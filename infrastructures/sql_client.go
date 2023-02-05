package infrastructures

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewSQLHandler() (*sql.DB, error) {
	handler, err := sql.Open("pgx", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=%s",
		"db",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_TIMEZONE"),
	))
	if err != nil {
		return nil, fmt.Errorf("incorrect DB connection information: %w", err)
	}

	if err := handler.Ping(); err != nil {
		return nil, fmt.Errorf("DB Ping failed: %w", err)
	}

	return handler, nil
}
