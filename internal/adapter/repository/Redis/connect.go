package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"marketflow/internal/adapter/config"
)

type Redis struct {
	Client *redis.Client
}

func Connect(ctx context.Context, cfg *config.App) (*Redis, error) {
	slog.Info("Connecting to Redis...")

	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Addr(),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.Database,
		PoolSize:     100,
		MinIdleConns: 100,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	slog.Info("Redis connected: %s", pong)
	return &Redis{Client: client}, nil
}
