package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golps/config"
	"golps/internal/app"
	"golps/pkg/logger"
)

func main() {
	srcFile := "app/main.go"

	cfg, err := config.New()
	if err != nil {
		fmt.Printf("%s    ERROR    %s:21  %v\n", curTime(), srcFile, err)
		os.Exit(1)
	}
	fmt.Printf("%s    INFO    %s:24  Config loaded\n", curTime(), srcFile)

	zapLogger, err := logger.New((*logger.Logger)(&cfg.Logger), cfg.App.Mode)
	if err != nil {
		fmt.Printf("%s    ERROR    %s:29  %v\n", curTime(), srcFile, err)
		os.Exit(1)
	}
	defer zapLogger.Sync()
	fmt.Printf("%s    INFO    %s:33  Logger initialized\n", curTime(), srcFile)

	ctx := context.Background()
	ctx = config.ToContext(ctx, cfg)
	ctx = logger.ToContext(ctx, zapLogger)

	zapLogger.Info("Application launching..")
	app.Launch(ctx)
}

func curTime() string {
	return time.Now().Format("2006-01-02T15:04:05.000Z0700")
}
