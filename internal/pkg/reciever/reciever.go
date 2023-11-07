package reciever

import (
	"errors"
	"homework-8/internal/pkg/kafka"

	"github.com/IBM/sarama"
)

type HandleFunc func(message *sarama.ConsumerMessage)

type KafkaReciever struct {
	consumer *kafka.Consumer
	handlers map[string]HandleFunc
}

func NewReceiver(consumer *kafka.Consumer, handlers map[string]HandleFunc) *KafkaReciever {
	return &KafkaReciever{
		consumer: consumer,
		handlers: handlers,
	}
}

func (r *KafkaReciever) Subscribe(topic string) error {
	handler, ok := r.handlers[topic]

	if !ok {
		return errors.New("can not find handler")
	}

	partitionList, err := r.consumer.SingleConsumer.Partitions(topic)

	if err != nil {
		return err
	}

	initialOffset := sarama.OffsetNewest

	for _, partition := range partitionList {
		pc, err := r.consumer.SingleConsumer.ConsumePartition(topic, partition, initialOffset)

		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			for message := range pc.Messages() {
				handler(message)
			}
		}(pc, partition)
	}

	return nil
}
