package tests

import (
	"homework-8/internal/pkg/kafka"
	test_consts "homework-8/tests/consts"
	"homework-8/tests/postgres"
	"log"
)

var (
	db   *postgres.TDB
	pr   *kafka.Producer
	cons *kafka.Consumer
)

func init() {
	db = postgres.NewFromEnv()
	var err error = nil
	pr, err = kafka.NewSyncProducer(test_consts.Brokers)
	cons, err = kafka.NewConsumer(test_consts.Brokers)
	if err != nil {
		log.Fatal("cant init kafka")
	}
}
