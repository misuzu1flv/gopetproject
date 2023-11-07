package main

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-8/internal/pkg/core"
	"homework-8/internal/pkg/db"
	"homework-8/internal/pkg/kafka"
	"homework-8/internal/pkg/logger"
	"homework-8/internal/pkg/reciever"
	"homework-8/internal/pkg/repository/postgresql"
	"homework-8/internal/pkg/sender"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
)

const port = ":9000"

var (
	brokers = []string{"127.0.0.1:9091", "127.0.0.1:9092", "127.0.0.1:9093"}
)

func main() {
	//DB init
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.GetPool(ctx).Close()

	//kafka producer init
	producer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}
	//kafka consumer init
	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	//init server
	implementation := core.NewServer(
		postgresql.NewPostRepo(db),
		postgresql.NewCommentRepo(db),
		sender.NewKafkaSender(producer, "logs"),
	)

	//init logger
	logger := logger.NewLogger(reciever.NewReceiver(consumer, map[string]reciever.HandleFunc{
		"logs": func(message *sarama.ConsumerMessage) {
			m := sender.Message{}
			err = json.Unmarshal(message.Value, &m)
			if err != nil {
				fmt.Println("Consumer error", err)
			}

			fmt.Println("Received Key: ", string(message.Key), " Value: ", m)
		},
	}))
	logger.StartLogger("logs")

	//start listening
	log.Println("Starting server")
	http.Handle("/", createRouter(implementation))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}

func createRouter(implementation *core.Server) *mux.Router {
	router := mux.NewRouter()

	router.Use(implementation.AddEventSender)

	router.HandleFunc("/post", func(w http.ResponseWriter, req *http.Request) {

		switch req.Method {

		case http.MethodGet:
			id, err := implementation.ParseId(req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			json, status := implementation.GetPostById(req.Context(), id)
			w.WriteHeader(status)
			w.Write(json)

		case http.MethodPost:

			body, err := implementation.ParsePost(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(implementation.CreatePost(req.Context(), body))

		case http.MethodDelete:

			log.Println("DeletePost")
			id, err := implementation.ParseId(req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(implementation.DeletePost(req.Context(), id))

		case http.MethodPatch:
			log.Println("UpdatePost")
			body, err := implementation.ParsePost(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(implementation.UpdatePost(req.Context(), body))

		default:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error bad method")
		}
	})

	router.HandleFunc("/comment", func(w http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodPost:
			body, err := implementation.ParseComment(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(implementation.CreateComment(req.Context(), body))
		default:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error bad method")
		}
	})

	return router
}
