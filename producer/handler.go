package main

import (
	"net/http"
	sarama "github.com/Shopify/sarama"
)

// KafkaController allows us to attach a producer to our handlers
type KafkaController struct {
	producer sarama.AsyncProducer
}

// Handler allows us to attach a KafkaController and send messages to the kafka producer queue asynchronously
func (kc *KafkaController) Handler(rw http.ResponseWriter, req *http.Request) {

	//Handle parsing error in message by returning a 400 http status code
	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	//Returns POST or PUT body parameters to msg
	msg := req.FormValue("msg")

	//Handle empty message by returning a 400 http status code
	if msg == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("msg can't be empty"))
		return
	}

	//A collection of elements are passed to the producer in order to send a message
	kc.producer.Input() <- &sarama.ProducerMessage{Topic: "example", Key: nil, Value: sarama.StringEncoder(msg)}

	//Returns a 202 status to the REST client when request is non empty or has no error during request parsing
	rw.WriteHeader(http.StatusAccepted)
}