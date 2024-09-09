package kafka

import (
	"context"
	"strings"

	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/sdk_errors"
)

type Consumer struct{}

func (Consumer) Subscribe(topics []string, chData chan []byte) error {
	kafkabrokers := os.Getenv("KAFKA_BROKERS")
	kafkaconsumergroup := os.Getenv("KAFKA_CONSUMER_GROUP")

	if kafkabrokers == "" || kafkaconsumergroup == "" {
		return sdk_errors.NewMissingInitialEnvironmentVariablesError("kafka", []string{"KAFKA_BROKERS", "KAFKA_CONSUMER_GROUP"})
	}
	keepRunning := true

	version, err := sarama.ParseKafkaVersion("3.0.0")
	if err != nil {
		return err
	}

	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = version

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := BaseConsumer{
		ready: make(chan bool),
		data:  chData,
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(kafkabrokers, ","), kafkaconsumergroup, consumerConfig)
	if err != nil {
		cancel()
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, topics, &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				fmt.Println(fmt.Sprintf("Error from consumer: %v", err))
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up

	// sigusr1 := make(chan os.Signal, 1)
	// signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			fmt.Println("terminating: context cancelled")
			keepRunning = false
			os.Exit(1)
		case <-sigterm:
			fmt.Println("terminating: via signal")
			keepRunning = false
			os.Exit(1)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		fmt.Println(fmt.Sprintf("Error closing client: %v", err))
		return err
	}
	return nil
}

// Consumer represents a Sarama consumer group consumer
type BaseConsumer struct {
	ready chan bool
	data  chan []byte
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *BaseConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *BaseConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *BaseConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Printf("message channel was closed")
				return nil
			}
			consumer.data <- message.Value
			session.MarkMessage(message, "")
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
