package kafka

import (
	"app/main/utils"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConfig struct {
	Host          string `json:"host"`
	ProducerTopic string `json:"producerTopic"`
	GroupId       string `json:"groupid"`
}

type Handler struct {
	producer *kafka.Producer
	config   *KafkaConfig
}

func InitKafkaHandler(jsonFileName string) *Handler {

	var handler Handler
	utils.ParseJsonConfig(jsonFileName, &handler.config)

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": handler.config.Host,
	})
	if err != nil {
		panic(err)
	}

	handler.producer = producer

	return &handler
}

func (h *Handler) StartKafkaHandler() {
	go h.startKafkaProducer()
}

func (h *Handler) Produce(message string) {
	h.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &h.config.ProducerTopic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}, nil)
}

func (h *Handler) startKafkaProducer() {

	go func() {
		for e := range h.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			case *kafka.Error:
				log.Printf("%% Error: %v\n", e)
			}
		}
	}()
}
