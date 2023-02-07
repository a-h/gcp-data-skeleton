package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type Client[T any] interface {
	Publish(ctx context.Context, msg Message[T]) (serverID string, err error)
}

func NewClient[T any](ctx context.Context, projectID, topicID string) (c Client[T], err error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		err = fmt.Errorf("pubsub.NewClient: %w", err)
	}
	c = &sender[T]{
		Topic: client.Topic(topicID),
	}
	return
}

type sender[T any] struct {
	Topic *pubsub.Topic
}

type Message[T any] struct {
	Data        T
	Attributes  map[string]string
	OrderingKey string
}

func (s sender[T]) Publish(ctx context.Context, msg Message[T]) (serverID string, err error) {
	data, err := json.Marshal(msg.Data)
	if err != nil {
		err = fmt.Errorf("pubsub.Publish: failed to marshal type to JSON: %w", err)
		return
	}
	return s.Topic.Publish(ctx, &pubsub.Message{
		Data:        data,
		Attributes:  msg.Attributes,
		OrderingKey: msg.OrderingKey,
	}).Get(ctx)
}
