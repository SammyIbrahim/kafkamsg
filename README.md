1. Install Zookeeper and Apache Kafka

https://www.tutorialspoint.com/apache_kafka/apache_kafka_basic_operations.htm

2. Consumer

In one terminal do the following..

cd kafkamsg/consumer
go run main.go


3. Producer

In another terminal do the following...

cd kafkamsg/producer
go run main.go


4. Use POSTMAN to send a request and receive response

Perform a POST request with http://localhost:3333 (will return a 202 status to POSTMAN when successful) and the following body format:

{
    "message": "this",
    "timestamp": "2021-05-23T23:46:00.00-04:00"
}

Note: Before running the producer or consumer, pull in the following dependencies from the cli while in the project root directory by typing the following:
1. go get github.com/Shopify/sarama
2. go get gopkg.in/go-playground/validator.v9

Some assumptions made:
1. Only one message at a time is to be sent asynchronously
