package kafka

import (
	"os"
	"strings"

	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/sdk_errors"
)

type kafkaProducer struct{}

func NewKafkaProducer() (KafkaProducer, error) {
	err := createProducer()
	if err != nil {
		return nil, err
	}
	return kafkaProducer{}, nil
}

var producer sarama.SyncProducer

func createProducer() error {
	kafkabrokers := os.Getenv("KAFKA_BROKERS")

	if kafkabrokers == "" {
		return sdk_errors.NewMissingInitialEnvironmentVariablesError("kafka", []string{"KAFKA_BROKERS"})
	}
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true
	_producer, err := sarama.NewSyncProducer(strings.Split(kafkabrokers, ","), producerConfig)
	if err != nil {
		return err
	}
	producer = _producer
	return nil
}

func (kp kafkaProducer) Send(topic string, eventType string, requestId string, producerName string, data []byte) error {

	message, err := kp.buildMessageFormat(eventType, requestId, producerName, data)
	if err != nil {
		return err
	}
	messageToProduce := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(message),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("requestId"),
				Value: []byte(requestId),
			},
			{
				Key:   []byte("eventType"),
				Value: []byte(eventType),
			},
		},
	}

	_, _, err = producer.SendMessage(messageToProduce)
	if err != nil {
		return err
	}

	return nil
}

func (kafkaProducer) buildMessageFormat(eventType string, requestId string, producerName string, data []byte) ([]byte, error) {
	response := fmt.Sprintf(`{"version":"1.3.0","requestId":"%s","timestamp":%d,"eventType":"%s","producer":"%s-withSdk","data":%s}`, requestId, time.Now().UnixNano()/int64(time.Second), eventType, producerName, string(data))
	return []byte(response), nil
}
