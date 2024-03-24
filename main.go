package main

import (
	"context"
	"fmt"
	"github.com/Andrmist/it-revolution-test-mine/internal"
	"github.com/Andrmist/it-revolution-test-mine/internal/domain"
	"github.com/Andrmist/it-revolution-test-mine/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := logrus.New()

	config, err := types.InitConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.PostgresHost, config.PostgresUser, config.PostgresPassword, config.PostgresDB, config.PostgresPort)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.Link{})

	serverCtx := types.ServerContext{
		Config:    config,
		Log:       logger,
		DB:        db,
		Validator: validator.New(),
	}

	internal.Run(ctx, serverCtx)

	serverCtx.Log.Info("Server is started!")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
	logger.Info("Shutdown...")
	cancel()
	os.Exit(0)
}
