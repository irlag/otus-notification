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
	Config                   *config.Config
	Log                      *zap.Logger
	Processors               *appProcessors.Processors
	Services                 *services.Services
	RecipeNotificationReader kafka.Reader[models.RecipeNotification]
}

func NewContainer(cfg *config.Config) *Container {
	logger, err := server.NewLogger(cfg.Debug)
	if err != nil {
		logger.Fatal("can't initialize zap logger", zap.Error(err))
	}

	kafkaNotificationWriter, _ := kafka.NewWriter[models.Notification](cfg.Kafka.Hosts)

	kafkaRecipeNotificationReader, _ := kafka.NewReader[models.RecipeNotification](
		cfg.Kafka.Hosts,
		models.NotificationConsumerGroup,
		models.RecipeEventKafkaTopic,
		func(notification models.RecipeNotification) {
			logger.Warn("error on consuming notification")
		},
	)

	srvs := services.New(logger, cfg)
	prcs := appProcessors.NewProcessor(kafkaNotificationWriter, logger)

	return &Container{
		Config:                   cfg,
		Log:                      logger,
		Processors:               prcs,
		Services:                 srvs,
		RecipeNotificationReader: kafkaRecipeNotificationReader,
	}
}
