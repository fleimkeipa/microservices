package consservice

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type ConsumerService struct {
	consumer sarama.Consumer
	topic    string
}

func NewConsumerService(consumer sarama.Consumer, topic string) *ConsumerService {
	return &ConsumerService{
		consumer: consumer,
		topic:    topic,
	}
}

func (rc *ConsumerService) Listen(topic string) {
	var msgCnt = 0

	consumer, err := rc.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("failed to partition consumer: %v", err)
		return
	}

	// Handle OS signals - used to stop the process.
	var sigchan = make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Create a Goroutine to run the consumer / worker.
	var doneCh = make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCnt++
				rc.handleMessage(msg.Value)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCnt, "messages")
}

func (rc *ConsumerService) handleMessage(data []byte) {
	// Process the message
	log.Printf("Processing message: %s", string(data))
}
