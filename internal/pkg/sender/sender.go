package sender

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type Producer interface {
	Close() error
	SendSyncMessage(message *sarama.ProducerMessage) (int32, int64, error)
}

type KafkaSender struct {
	producer Producer
	topic    string
}

type Message struct {
	EventType string    `json:"eventType"`
	Request   string    `json:"request"`
	Time      time.Time `json:"time"`
}

func NewKafkaSender(pr Producer, t string) *KafkaSender {
	return &KafkaSender{producer: pr, topic: t}
}

func (k *KafkaSender) SendMessage(message Message) error {
	producerMessage, err := k.buildMessage(message)
	if err != nil {
		return err
	}

	_, _, err = k.producer.SendSyncMessage(producerMessage)
	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaSender) buildMessage(message Message) (*sarama.ProducerMessage, error) {
	marsh, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return &sarama.ProducerMessage{
		Topic: k.topic,
		Key:   sarama.StringEncoder(fmt.Sprint(message.EventType)),
		Value: sarama.ByteEncoder(marsh),
	}, nil
}
