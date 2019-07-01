package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/esvm/dcm-middleware/dcm"
	"github.com/esvm/dcm-middleware/distribution/models"
	"github.com/esvm/dcm-middleware/scylla"
	"github.com/esvm/microservices/src/database_client"
	"github.com/esvm/microservices/src/proto"
)

func main() {
	host := os.Getenv("BROKER_HOST")
	string_port := os.Getenv("BROKER_PORT")
	port, err := strconv.Atoi(string_port)
	if err != nil {
		fmt.Println(err)
		return
	}

	grpcAddress := os.Getenv("GRPC_ADDRESS")
	grpcClient, err := database_client.NewGrpcDatabaseClient(grpcAddress)
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
		consume(c, topic, grpcClient)
	}
}

func consume(c *dcm.Connection, topic *models.Topic, grpcClient database_client.DatabaseClient) {
	ch, err := c.Consume(topic.ID, topic.Properties.IndexName)
	if err != nil {
		fmt.Println(err)
	}

	for metric := range ch {
		fmt.Println(metric)

		err := grpcClient.InsertItem(context.Background(), &proto.InsertItemRequest{
			TopicName: "",
			Metrics: []*proto.Metric{
				&proto.Metric{
					Average:           float32(metric.Average),
					Median:            float32(metric.Median),
					Variance:          float32(metric.Variance),
					StandardDeviation: float32(metric.StandardDeviation),
					Mode:              float32(metric.Mode),
				},
			},
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}
