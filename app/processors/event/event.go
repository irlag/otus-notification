package event

import (
	"context"

	"otus-notification/app/api/parameters"
)

type Event interface {
	Send(ctx context.Context, params *parameters.EventSendParams) error
}

type eventProcessor struct{}

func NewEventProcessor() Event {
	return &eventProcessor{}
}
