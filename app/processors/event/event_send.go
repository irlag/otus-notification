package event

import (
	"context"

	"otus-notification/app/api/parameters"
)

func (r *eventProcessor) Send(ctx context.Context, params *parameters.EventSendParams) error {
	return nil
}
