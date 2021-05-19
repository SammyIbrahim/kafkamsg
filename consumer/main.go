package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	//Configure a new Sarama Consumer with given broker Addresses and configuration
	SaramaConsumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer SaramaConsumer.Close()

	//Create a consumer with the given topic and offset
	partitionConsumer, err := SaramaConsumer.ConsumePartition("example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	//Consumer continuously polls for messages returned by the broker and outputs to console
	for {
		consumerMsg := <-partitionConsumer.Messages()
		log.Printf("Message from Sarama queue: \"%s\" at offset: %d\n", consumerMsg.Value, consumerMsg.Offset)
	}
}
