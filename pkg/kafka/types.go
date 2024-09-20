package kafka

type KafkaProducer interface {
	// Send sends a message to a kafka topic
	// topic: the topic to send the message to
	// eventType: the type of event that is being sent
	// requestId: the id of the request that generated the event
	// producerName: the name of the producer that is sending the message
	// data: the data to send
	Send(topic string, eventType string, requestId string, producerName string, data []byte) error
}
