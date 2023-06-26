package app

import (
	"go.uber.org/zap"

	"otus-notification/app/models"
	appProcessors "otus-notification/app/processors"
	"otus-notification/app/services"
	"otus-notification/app/storage/kafka"

	"otus-notification/app/config"
	"otus-notification/app/server"
)

type Container struct {
	Config     *config.Config
	Log        *zap.Logger
	Processors *appProcessors.Processors
	Services   *services.Services
}

func NewContainer(cfg *config.Config) *Container {
	logger, err := server.NewLogger(cfg.Debug)
	if err != nil {
		logger.Fatal("can't initialize zap logger", zap.Error(err))
	}

	kafkaRecipeWriter, _ := kafka.NewWriter[models.Notification](cfg.Kafka.ZookeeperHosts)
	srvs := services.New(logger, cfg)
	prcs := appProcessors.NewProcessor(kafkaRecipeWriter)

	return &Container{
		Config:     cfg,
		Log:        logger,
		Processors: prcs,
		Services:   srvs,
	}
}
