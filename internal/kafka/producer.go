package kafka

import (
	"log"

	"github.com/akolybelnikov/xm-exercise/internal/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MutationProducer interface {
	Produce(topic string, key string, value string) error
	Errors()
}

type Producer struct {
	producer *kafka.Producer
	errors   chan error
	delivery chan kafka.Event
}

func NewProducer(cfg *config.KafkaConfig) (*Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
	})
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: producer,
		errors:   make(chan error),
		delivery: make(chan kafka.Event, cfg.ChanSize),
	}, nil
}

func (p *Producer) Start() {
	go p.handleDeliveries()
}

func (p *Producer) Produce(topic string, key string, value string) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}, p.delivery)
}

func (p *Producer) handleDeliveries() {
	for e := range p.delivery {
		m, ok := e.(*kafka.Message)
		if !ok {
			log.Println("Ignored event: ", e)
			continue
		}
		if m.TopicPartition.Error != nil {
			p.errors <- m.TopicPartition.Error
		}
	}
}

func (p *Producer) Errors() {
	for err := range p.errors {
		log.Printf("Kafka error: %v", err)
	}
}

func (p *Producer) Close(timeout int) {
	log.Println("Closing producer")
	p.producer.Flush(timeout)
	p.producer.Close()
	close(p.errors)
	close(p.delivery)
}
