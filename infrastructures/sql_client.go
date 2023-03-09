package infrastructures

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewSQLHandler() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s Timezone=%s",
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)
	if _, err := os.Stat("root.crt"); !errors.Is(err, os.ErrNotExist) {
		// 証明書が存在しているとき
		dsn = fmt.Sprintf("%s sslmode=verify-full sslrootcert=root.crt", dsn)
		log.Println("sslmode=verify-full")
	} else {
		// 証明書がないとき
		dsn = fmt.Sprintf("%s sslmode=disable", dsn)
		log.Println("sslmode=disable")
	}

	handler, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("incorrect DB connection information: %w", err)
	}

	if err := handler.Ping(); err != nil {
		return nil, fmt.Errorf("DB Ping failed: %w", err)
	}

	return handler, nil
}
