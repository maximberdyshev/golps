package app

import (
	"context"

	"golps/config"
	"golps/internal/usecase"
	"golps/internal/usecase/repo"
	"golps/internal/usecase/webapi"
	"golps/pkg/logger"
	"golps/pkg/redis"

	"go.uber.org/zap"
)

func Launch(ctx context.Context) {
	logger := logger.FromContext(ctx)
	cfg := config.FromContext(ctx)

	redis, err := redis.New(cfg.Redis.Path)
	if err != nil {
		logger.Fatal("redis", zap.Error(err))
	}

	repo, err := repo.New(ctx, redis)
	if err != nil {
		logger.Fatal("repo", zap.Error(err))
	}

	webAPI, err := webapi.New(ctx)
	if err != nil {
		logger.Fatal("webapi", zap.Error(err))
	}

	U := usecase.New(ctx, repo, webAPI)
	U.Start()
}
