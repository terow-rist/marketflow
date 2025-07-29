package PostgreSQL

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log/slog"
	"marketflow/internal/adapter/config"
	"time"
)

type PostgreSQL struct {
	db *sql.DB
}

func Connect(ctx context.Context, config *config.App) (*PostgreSQL, error) {
	slog.Info("Connecting to database...")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	slog.Info("Connected to database")
	return &PostgreSQL{db: db}, err
}
