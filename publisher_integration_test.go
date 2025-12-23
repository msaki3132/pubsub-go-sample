package main

import (
	"context"
	"os"
	"testing"
	"time"

	pubsubpb "example.com/pubsub-go-sample/proto"
	"github.com/google/uuid"
	// 同じくパス調整
)

func TestPublishToRealTopic_Integration(t *testing.T) {
	// -count=1 でキャッシュ無効化推奨
	// go test -count=1 -run Integration -v

	ctx := context.Background()

	projectID := os.Getenv("GCP_PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	if projectID == "" || topicName == "" {
		t.Fatal("GCP_PROJECT_ID and PUBSUB_TOPIC are required")
	}

	// protobufメッセージ作成（毎回違う内容に）
	testMsg := &pubsubpb.TestMessage{
		Greeting:  "hello from protobuf integration test",
		Timestamp: time.Now().UnixMilli(),
		TestId:    uuid.New().String(), // 確実に一意
	}

	msgID, err := PublishMessage(ctx, projectID, topicName, testMsg)
	if err != nil {
		t.Fatalf("PublishMessage failed: %v", err)
	}

	t.Logf("Successfully published protobuf message")
	t.Logf("Message ID: %s", msgID)
	t.Logf("Content: greeting=%s, timestamp=%d, test_id=%s",
		testMsg.Greeting, testMsg.Timestamp, testMsg.TestId)
}
