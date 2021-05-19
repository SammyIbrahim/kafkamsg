package main

import (
	"github.com/Shopify/sarama"
	"log"
)

// GetResponse grabs results and errors from a producer asynchronously
func GetResponse(asyncProducer sarama.AsyncProducer) {
	for {
		select {
		case res := <-asyncProducer.Successes():
			log.Printf("> The message: \"%s\" was sent to partition %d at offset %d\n", res.Value, res.Partition, res.Offset)
		case err := <-asyncProducer.Errors():
			log.Println("Error creating message", err)
		}
	}
}
