package main

import (
	"fmt"

	"github.com/esvm/dcm-middleware/dcm"
	"github.com/esvm/dcm-middleware/distribution/models"
	"github.com/esvm/dcm-middleware/scylla"
)

const host = "localhost"

func main() {
	c, err := dcm.Connect(host, 8426, 1)
	defer c.Close()
	topic, err := c.CreateTopic("topic_name", models.TopicProperties{IndexName: "topic_name_index", StartFrom: scylla.Begin})
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 2; i++ {
		publish(c, topic)
	}
}

func publish(c *dcm.Connection, topic *models.Topic) {
	message := &models.Message{
		TopicID: topic.ID,
		Payload: 1.5,
	}

	err := c.Publish(topic.ID, message)
	if err != nil {
		fmt.Println(err)
	}
}
