package kafka

import "github.com/segmentio/kafka-go"

func getWriter(kafkaURL, topic string) *kafka.Writer {
	kw := new(kafka.Writer)
	kw.Addr = kafka.TCP(kafkaURL)
	kw.Topic = topic
	kw.Balancer = &kafka.LeastBytes{} // Default load balancer
	return kw
}
