package main

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/proto"

	pubsubpb "example.com/pubsub-go-sample/proto"
)

func PublishMessage(ctx context.Context, projectID, topicID string, msg *pubsubpb.TestMessage) (string, error) {
	// protobuf をバイト列にマーシャル
	data, err := proto.Marshal(msg)
	if err != nil {
		return "", err
	}

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return "", err
	}
	defer client.Close()

	topic := client.Topic(topicID)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
		Attributes: map[string]string{
			"content_type": "application/protobuf",
			"message_type": "pubsub.TestMessage",
			"sent_at":      time.Now().UTC().Format(time.RFC3339),
		},
	})

	msgID, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return msgID, nil
}
