package processors

import (
	"otus-notification/app/processors/event"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
	EventProcessor       event.Event
}

func NewProcessor() *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
		EventProcessor:       event.NewEventProcessor(),
	}
}
