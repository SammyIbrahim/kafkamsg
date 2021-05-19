package main

import (
	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestKafkaController_Handler(t *testing.T) {
	//Returns a mock controller
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	//Returns a mock of an Async Producer
	mockAsyncProducer := NewMockAsyncProducer(mockController)

	type fields struct {
		producer sarama.AsyncProducer
	}

	type args struct {
		rw http.ResponseWriter
		req *http.Request
	}

	//Test object
	tests := []struct {
		name   string
		fields fields
		arguments args
	}{
		{"empty message", fields{mockAsyncProducer}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
	}

	//Runs test for KafkaController Handler
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kc := &KafkaController{
				producer: tt.fields.producer,
			}
			kc.Handler(tt.arguments.rw, tt.arguments.req)
		})
	}
}