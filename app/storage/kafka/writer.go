package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"

	"otus-notification/app/models"
)

type Writer[T models.Notification] struct {
	writer *kafka.Writer
}

func NewWriter[T models.Notification](hosts string) (Writer[T], func() error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(hosts),
		Balancer: &kafka.LeastBytes{},
	}

	return Writer[T]{writer: w}, w.Close
}

func (w *Writer[T]) Write(ctx context.Context, item T) error {
	m, err := item.MarshalJSON()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Value: m,
		Topic: item.GetEventTopic(),
	}

	return w.writer.WriteMessages(ctx, message)
}
