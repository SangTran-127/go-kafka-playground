package kafka

import (
	"github.com/segmentio/kafka-go"
)

var (
	kafkaConsumer *kafka.Writer
)

const (
	kafkaURL = "local"
)
