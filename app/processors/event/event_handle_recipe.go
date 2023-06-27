package event

import (
	"fmt"

	"otus-notification/app/models"
)

func (r *eventProcessor) HandleRecipeNotification(recipe models.RecipeNotification) error {
	r.log.Info(fmt.Sprintf("recived recipe message %s", recipe.Name))

	return nil
}
