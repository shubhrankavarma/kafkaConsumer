package main

import (
	"fmt"

	kafkaConfig "github.com/amagimedia/kafkaConsumer/kafka"
)

func main() {
	fmt.Println("Started Service")

	// kafkaConfig.StartKafka()
	kafkaConfig.StartKafka()

}
