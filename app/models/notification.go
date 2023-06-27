package models

import "encoding/json"

const (
	NotificationConsumerGroup = "otus.notification"
)

type Notification interface {
	json.Marshaler
	GetEventTopic() string
}
