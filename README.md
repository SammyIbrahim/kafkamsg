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


4. Use POSTMAN AND curl to send a request and receive response

For Postman, a POST request with: http://localhost:3333/?msg="this" (will show a 202 status returned)

OR

in a terminal..
curl "http://localhost:3333/?msg="this" (will show the message returned to the consumer's console)

