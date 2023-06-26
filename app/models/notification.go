package models

type Notification interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	GetEventTopic() string
}
