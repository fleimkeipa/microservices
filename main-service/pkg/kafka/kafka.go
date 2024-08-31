package kafka

import "github.com/IBM/sarama"

func ConnectToKafka() (sarama.SyncProducer, error) {
	var config = sarama.NewConfig()
	config.Producer.Return.Successes = true

	var brokers = []string{"kafka:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
