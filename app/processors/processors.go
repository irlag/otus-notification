package processors

import (
	"otus-notification/app/models"
	"otus-notification/app/processors/event"
	kafkaStorage "otus-notification/app/storage/kafka"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
	EventProcessor       event.Event
}

func NewProcessor(kafkaWriter kafkaStorage.Writer[models.Notification]) *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
		EventProcessor:       event.NewEventProcessor(kafkaWriter),
	}
}
