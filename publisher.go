package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/esvm/dcm-middleware/dcm"
	"github.com/esvm/dcm-middleware/distribution/models"
	"github.com/esvm/dcm-middleware/scylla"
)

var (
	diskUsedTopic *models.Topic
)

const host = "localhost"

type data struct {
	diskUsed float64
}

func getInfoHW() data {
	time.Sleep(500 * time.Millisecond)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	number := r1.Float64() * 100
	fmt.Println(number)
	return data{
		diskUsed: number,
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

func createTopics(c *dcm.Connection) error {
	var err error

	diskUsedTopic, err = c.CreateTopic("disk_used", models.TopicProperties{IndexName: "disk_used_index", StartFrom: scylla.Begin})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	c, err := dcm.Connect(host, 8426, 10)
	defer c.Close()

	err = createTopics(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		dataHW := getInfoHW()
		publish(c, diskUsedTopic, dataHW.diskUsed)
	}
}
