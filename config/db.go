package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnectionString(
	cfg *EnvConfig,
) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresName,
	)
}

func NewDBConnection(
	ctx context.Context,
	cfg EnvConfig,
) (*pgxpool.Pool, error) {
	connString := CreateConnectionString(&cfg)

	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
