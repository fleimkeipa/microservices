package repositories

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type KafkaRepository struct {
	producer sarama.SyncProducer
}

func NewKafkaRepository(producer sarama.SyncProducer) *KafkaRepository {
	return &KafkaRepository{
		producer: producer,
	}
}

func (rc *KafkaRepository) Send(subj string, data string) error {
	var msg = sarama.ProducerMessage{
		Topic: subj,
		Value: sarama.StringEncoder(data),
	}

	// Create order
	partition, offset, err := rc.producer.SendMessage(&msg)
	if err != nil {
		return fmt.Errorf("failed to publish, topic: [%s], data: [%s], error: %w", subj, data, err)
	}

	log.Printf("order is stored in topic: [%s], partition: [%d], offset: [%d]", subj, partition, offset)

	return nil
}
