package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"otus-notification/app/models"
)

type Reader[T models.Notification] struct {
	reader  *kafka.Reader
	onError func(item T)
}

func NewReader[T models.Notification](hosts, group, topic string, onError func(T)) (Reader[T], func() error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{hosts},
		Topic:   topic,
		GroupID: group,
	})

	return Reader[T]{reader: r, onError: onError}, r.Close
}

func (r Reader[T]) Read(handler func(item T) error) error {
	for {
		message, err := r.reader.FetchMessage(context.TODO())
		if err != nil {
			return err
		}

		var t T
		err = json.Unmarshal(message.Value, &t)
		if err != nil {
			r.onError(t)
		}

		err = handler(t)
		if err != nil {
			r.onError(t)
		}

		err = r.reader.CommitMessages(context.TODO(), message)
		if err != nil {
			return err
		}
	}
}
