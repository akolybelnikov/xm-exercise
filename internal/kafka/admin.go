package kafka

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func CreateTopic(broker, topic string) error {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		return err
	}
	defer adminClient.Close()

	ctx := context.Background()
	results, err := adminClient.CreateTopics(ctx, []kafka.TopicSpecification{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	})

	log.Println("Results:", results)

	return err
}
