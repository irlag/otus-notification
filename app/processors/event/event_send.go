package event

import (
	"context"
	"errors"

	"otus-notification/app/api/parameters"
	"otus-notification/app/models"
)

func (r *eventProcessor) Send(ctx context.Context, params *parameters.EventSendParams) error {
	if params.Name == models.RecipeEventName {
		recipeNotification := &models.RecipeNotification{}
		err := recipeNotification.UnmarshalJSON([]byte(params.Data))
		if err != nil {
			return err
		}

		err = r.kafkaWriter.Write(ctx, recipeNotification)
		if err != nil {
			return err
		}
	} else {
		return errors.New("event name is incorrect")
	}

	return nil
}
