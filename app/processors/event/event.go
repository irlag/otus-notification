package event

import (
	"context"

	"go.uber.org/zap"

	"otus-notification/app/api/parameters"
	"otus-notification/app/models"
	"otus-notification/app/storage/kafka"
)

type Event interface {
	Send(ctx context.Context, params *parameters.EventSendParams) error
	HandleRecipeNotification(recipe models.RecipeNotification) error
}

type eventProcessor struct {
	kafkaWriter kafka.Writer[models.Notification]
	log         *zap.Logger
}

func NewEventProcessor(
	kafkaWriter kafka.Writer[models.Notification],
	logger *zap.Logger,
) Event {
	return &eventProcessor{kafkaWriter: kafkaWriter, log: logger}
}
