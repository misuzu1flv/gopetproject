package loggertest

import (
	"encoding/json"
	"fmt"
	"homework-8/internal/pkg/kafka"
	"homework-8/internal/pkg/logger"
	"homework-8/internal/pkg/reciever"
	"homework-8/internal/pkg/sender"

	"github.com/IBM/sarama"
)

func SetUpTest(cons *kafka.Consumer) chan string {
	resultch := make(chan string)
	logger := logger.NewLogger(reciever.NewReceiver(cons, map[string]reciever.HandleFunc{
		"logs": func(message *sarama.ConsumerMessage) {
			m := sender.Message{}
			err := json.Unmarshal(message.Value, &m)
			if err != nil {
				fmt.Println("Consumer error", err)
			}

			resultch <- fmt.Sprint(string(message.Key))
		},
	}))
	logger.StartLogger("logs")
	return resultch
}
