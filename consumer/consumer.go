package main

import (
	"fmt"
	"os"
	"strconv"

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

	c, err := dcm.Connect(host, port, 100)
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
		consume(c, topic)
	}
}

func consume(c *dcm.Connection, topic *models.Topic) {
	ch, err := c.Consume(topic.ID, topic.Properties.IndexName)
	if err != nil {
		fmt.Println(err)
	}

	for metric := range ch {
		fmt.Println(metric)
	}
}
