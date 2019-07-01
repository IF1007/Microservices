package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"

	"github.com/esvm/dcm-middleware/dcm"
	"github.com/esvm/dcm-middleware/distribution/models"
	"github.com/esvm/dcm-middleware/scylla"
)

func main() {
	host := os.Getenv("BROKER_HOST")
	string_port := os.Getenv("BROKER_PORT")
	port, err := strconv.Atoi(string_port)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := dcm.Connect(host, port, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	topic_name := os.Getenv("BROKER_TOPIC_NAME")
	topic_name_index := os.Getenv("BROKER_TOPIC_INDEX")
	topic, err := c.CreateTopic(topic_name, models.TopicProperties{IndexName: topic_name_index, StartFrom: scylla.Begin})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		publish(c, topic, rand.Float64()*2)
	}
}

func publish(c *dcm.Connection, topic *models.Topic, payload float64) {
	message := &models.Message{
		TopicID: topic.ID,
		Payload: payload,
	}

	err := c.Publish(topic.ID, message)
	if err != nil {
		fmt.Println(err)
	}
}
