package kafka_test

import (
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/config"
	"github.com/akolybelnikov/xm-exercise/internal/kafka"
	"github.com/stretchr/testify/require"
)

func TestNewProducer(t *testing.T) {
	testCases := []struct {
		name      string
		config    *config.KafkaConfig
		expectErr bool
	}{
		{
			name:      "invalid_config",
			config:    &config.KafkaConfig{},
			expectErr: false,
		},
		{
			name: "valid_config",
			config: &config.KafkaConfig{
				Brokers:  "localhost:9092",
				ChanSize: 10,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := kafka.NewMutationProducer(tc.config)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, ret)
			}
		})
	}
}

func TestProducer_Produce(t *testing.T) {
	producer, _ := kafka.NewMutationProducer(&config.KafkaConfig{
		Brokers:  "localhost:9092",
		ChanSize: 1,
	})

	err := producer.Produce("test", "key", "value")

	require.NoError(t, err)
}

func TestProducer_Start(t *testing.T) {
	producer, _ := kafka.NewMutationProducer(&config.KafkaConfig{
		Brokers:  "localhost:9092",
		ChanSize: 1,
	})

	producer.Start()

	t.Log("Producer started")
}

func TestProducer_Close(t *testing.T) {
	producer, _ := kafka.NewMutationProducer(&config.KafkaConfig{
		Brokers:  "localhost:9092",
		ChanSize: 1,
	})
	producer.Start()
	producer.Close(1000)

	t.Log("Producer closed")
}

// func TestKafkaProducer(t *testing.T) {
//	p, err := kafka.NewMutationProducer(&config.KafkaConfig{
//		Brokers:  "localhost:29092",
//		ChanSize: 1,
//	})
//	require.NoError(t, err)
//
//	defer func() {
//		t.Log("Closing producer")
//		p.Flush(1000)
//		p.Close(1000)
//	}()
//
//	err = p.Produce("test", "key", "value")
//	require.NoError(t, err)
//
//	// Wait for the message to be delivered
//	e := <-p.Events()
//	m, ok := e.(*kafka2.Message)
//	if !ok {
//		t.Fatalf("Ignored event: %v", e)
//	}
//
//	if m.TopicPartition.Error != nil {
//		t.Errorf("Message delivery failed: %v", m.TopicPartition.Error)
//	} else {
//		t.Logf("Message delivered to topic %s [%d] at offset %v",
//			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
//	}
//}
