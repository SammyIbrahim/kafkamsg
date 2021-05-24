package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// KafkaController allows us to attach a producer to our handlers
type KafkaController struct {
	producer sarama.AsyncProducer
}

//Defines the struct within which the json request will be UnMarshalled
type MessageRequest struct {
	Message string `json:"message" validate:"required"`
	TimeStamp string `json:"timeStamp" validate:"required"`
}

// Handler allows us to attach a KafkaController and send messages to the kafka producer queue asynchronously
func (kc *KafkaController) Handler(rw http.ResponseWriter, req *http.Request) {

	//Creates a new validator instance to be used to validate that the request body has both required fields
	validate := validator.New()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(req.Body)

	//Read Request Body
	b, reqErr := ioutil.ReadAll(req.Body)

	//UnMarshall MessageRequest struct
	var messageRequest MessageRequest
	jsonErr := json.Unmarshal(b, &messageRequest)

	//Validator will validate that the messageRequest struct has both the required fields
	vErr := validate.Struct(messageRequest)

	//Extract message and time values from request
	msg := messageRequest.Message
	ts := messageRequest.TimeStamp

	//Formats the timestamp into RFC3339Nano format
	tsNano, parseErr := time.Parse(time.RFC3339Nano, ts)

	//Timestamp with value of current time
	now := time.Now()

	//Time difference between now and nano timestamp generated from request
	diff := tsNano.Sub(now)

	//Error handling cases
	switch {

	//Handle other errors that occur during request body reading or json UnMarshalling
	case (reqErr != nil) || (jsonErr != nil):
		rw.WriteHeader(http.StatusInternalServerError)
		return

	//Case where both required fields are not specified in request body
	case vErr !=nil:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("The message and timestamp fields are both required. Please check your request."))
		return

	//Case of time parsing error
	case parseErr !=nil:
		rw.WriteHeader(http.StatusBadRequest)
		return

	//Case where messaging time requested occurs in the past
	case diff < 0:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Messaging time cannot occur in the past"))
		return
	}

	//AfterFunc waits for the duration to elapse and then calls func in it's own go routine
	time.AfterFunc(diff, func(){
		//A collection of elements are passed to the producer in order to send a message
		kc.producer.Input() <- &sarama.ProducerMessage{Topic: "example", Key: nil, Value: sarama.StringEncoder(msg)}
	})

	//Returns a 202 status to the REST client when request is non empty or has no error during request parsing
	rw.WriteHeader(http.StatusAccepted)
}