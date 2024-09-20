package repo

import (
	"context"

	"golps/pkg/redis"
)

// TODO: placeholder

type Repo struct{}

func New(ctx context.Context, redis *redis.Redis) (*Repo, error) {
	return &Repo{}, nil
}
