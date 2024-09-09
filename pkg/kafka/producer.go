package kafka

import (
	"os"
	"strings"

	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/Netsocs-Team/netsocs_go_sdk/pkg/sdk_errors"
)

type KafkaProducer struct{}

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
		panic(err)
	}
	producer = _producer
	return nil
}

func GetKafkaProducer() (sarama.SyncProducer, error) {
	if producer == nil {
		err := createProducer()
		if err != nil {
			return nil, err
		}

	}
	return producer, nil
}

func (KafkaProducer) BuildMessageFormat(eventType string, requestId string, producer string, data []byte) ([]byte, error) {
	response := fmt.Sprintf(`{"version":"1.3.0","requestId":"%s","timestamp":%d,"eventType":"%s","producer":"%s","data":%s}`, requestId, time.Now().UnixNano()/int64(time.Second), eventType, producer, string(data))
	return []byte(response), nil
}
