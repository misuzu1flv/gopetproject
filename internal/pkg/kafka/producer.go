package kafka

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	brokers      []string
	syncProducer sarama.SyncProducer
}

func NewSyncProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	saramaProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {

		return nil, err
	}

	Producer := &Producer{brokers: brokers, syncProducer: saramaProducer}
	return Producer, nil
}

func (k *Producer) SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return k.syncProducer.SendMessage(message)
}

func (k *Producer) Close() error {
	err := k.syncProducer.Close()
	return err
}
