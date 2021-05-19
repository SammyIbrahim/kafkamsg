package main

import (
	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetResponse(t *testing.T) {
	//Returns a mock controller
	controller := gomock.NewController(t)
	defer controller.Finish()

	//Returns a mock of an Async Producer
	mockAsyncProducer := NewMockAsyncProducer(controller)

	type args struct {
		producer sarama.AsyncProducer
	}

	//Test object
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{mockAsyncProducer}},
	}

	//Runs test for Producer
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := make(chan *sarama.ProducerError)
			success := make(chan *sarama.ProducerMessage)

			mockAsyncProducer.EXPECT().Successes().AnyTimes().Return(success)
			mockAsyncProducer.EXPECT().Errors().AnyTimes().Return(errors)

			go GetResponse(tt.args.producer)

			//When a producer fails to deliver a message an error is sent to errors channel
			errors <- &sarama.ProducerError{}

			//When a producer delivers a message it is sent to the success channel
			success <- &sarama.ProducerMessage{}
		})
	}
}