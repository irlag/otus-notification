package event

import (
	"context"

	"otus-notification/app/api/parameters"
	"otus-notification/app/models"
	"otus-notification/app/storage/kafka"
)

type Event interface {
	Send(ctx context.Context, params *parameters.EventSendParams) error
}

type eventProcessor struct {
	kafkaWriter kafka.Writer[models.Notification]
}

func NewEventProcessor(kafkaWriter kafka.Writer[models.Notification]) Event {
	return &eventProcessor{kafkaWriter: kafkaWriter}
}
