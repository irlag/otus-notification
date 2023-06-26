package models

//go:generate easyjson

const (
	RecipeEventKafkaTopic = "otus.recipe.update"
	RecipeEventName       = "recipe.notification"
)

//easyjson:json
type RecipeNotification struct {
	Name string `json:"name"`
}

func (r RecipeNotification) GetEventTopic() string {
	return RecipeEventKafkaTopic
}
