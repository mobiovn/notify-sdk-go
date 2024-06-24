package notifysdk

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func pushMessageToTopicReceive(data map[string]interface{}) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers":  Env.KafkaBroker,
		"compression.type":   "zstd",
		"request.timeout.ms": 20000,
	}

	producer, err := kafka.NewProducer(conf)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	doneChan := make(chan bool)

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully delivered message to %v\n", ev.TopicPartition)
				}
				doneChan <- true
			}
		}
	}()

	topicSDK := Topic.TopicNotifySdkReceiveMessage
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicSDK, Partition: kafka.PartitionAny},
		Value:          dataBytes,
	}, nil)
	<-doneChan

	if err != nil {
		panic(err)
	}
	close(doneChan)
	fmt.Printf("[NOTIFY_SDK]: push_message_to_topic: topic: %s\n", topicSDK)
	fmt.Printf("[NOTIFY_SDK]: push_message_to_topic: message: %s\n", string(dataBytes))
}
