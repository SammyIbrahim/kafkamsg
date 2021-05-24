package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"net/http"
)

func main() {
	//Configure sarama Producer struct with sane defaults
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true

	//cron := cron.Schedule(time.Minute)

	//Configure a new Async Producer with given broker Addresses and configuration
	asyncProducer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		panic(err)
	}
	defer asyncProducer.AsyncClose()

	//GetResponse will grab results and errors from a Sarama Producer Asynchronously
	go GetResponse(asyncProducer)

	//Allows us to attach a producer to our handlers
	c := KafkaController{asyncProducer}

	//Registers a handler for a given pattern
	http.HandleFunc("/", c.Handler)

	fmt.Println("Listening on port :3333")
	panic(http.ListenAndServe(":3333", nil))
}
