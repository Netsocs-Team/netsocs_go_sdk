package kafka

import (
	"encoding/json"
	"fmt"
)

type KafkaMessageSchema struct {
	Version   string `json:"version"`
	RequestID string `json:"requestId"`
	Timestamp int64  `json:"timestamp"`
	EventType string `json:"eventType"`
	Producer  string `json:"producer"`
	Data      []byte `json:"data"`
}

// this struct is used to unmarshal the kafka message
type preKafkaMessageSchema struct {
	Version   string      `json:"version"`
	RequestID string      `json:"requestId"`
	Timestamp int64       `json:"timestamp"`
	EventType string      `json:"eventType"`
	Producer  string      `json:"producer"`
	Data      interface{} `json:"data"`
}

type ErrUnmarshalKafkaMessage struct {
	Received string
}

func (e ErrUnmarshalKafkaMessage) Error() string {
	return fmt.Sprintf("Error unmarshalling kafka message, expected `{'version': 'value', 'requestId': 'value', 'timestamp': 'value', 'eventType': 'value', 'producer': 'value', 'data': 'value'}`, got: %s", e.Received)
}

func UnmarshalKafkaMessageSchema(data []byte) (KafkaMessageSchema, error) {
	r := preKafkaMessageSchema{}

	err := json.Unmarshal(data, &r)
	if err != nil {
		return KafkaMessageSchema{}, err
	}
	kafkaMessageSchemaData, err := json.Marshal(r.Data)
	if err != nil {
		return KafkaMessageSchema{}, err
	}
	return KafkaMessageSchema{
		Version:   r.Version,
		RequestID: r.RequestID,
		Timestamp: r.Timestamp,
		EventType: r.EventType,
		Producer:  r.Producer,
		Data:      kafkaMessageSchemaData,
	}, nil
}
