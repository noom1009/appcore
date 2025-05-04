package messagequeue

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func InitKafka(broker, topic string) {
    KafkaWriter = &kafka.Writer{
        Addr:     kafka.TCP(broker),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
    }
    log.Println("✅ Kafka initialized on broker:", broker)
}

func SendMessage(key, value string) error {
    msg := kafka.Message{
        Key:   []byte(key),
        Value: []byte(value),
    }
    return KafkaWriter.WriteMessages(context.Background(), msg)
}
