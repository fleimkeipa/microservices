package kafka

import "github.com/IBM/sarama"

func ConnectConsumer() (sarama.Consumer, error) {
	var brokers = []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
