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
	stringPort := os.Getenv("BROKER_PORT")
	port, err := strconv.Atoi(stringPort)
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

	topicName := os.Getenv("BROKER_TOPIC_NAME")
	topicNameIndex := os.Getenv("BROKER_TOPIC_INDEX")
	topic, err := c.CreateTopic(topicName, models.TopicProperties{IndexName: topicNameIndex, StartFrom: scylla.End})
	if err != nil {
		fmt.Println(err)
		return
	}

	consume(c, topic, grpcClient)
}

func consume(c *dcm.Connection, topic *models.Topic, grpcClient database_client.DatabaseClient) {
	ch, err := c.Consume(topic.ID, topic.Properties.IndexName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("AQUIAQUIAQUIAQUIAQUIAQUIAQUIAQUIAQUI")
	fmt.Println(topic)
	fmt.Println("AQUIAQUIAQUIAQUIAQUIAQUIAQUIAQUIAQUI")

	for metric := range ch {
		// fmt.Println(metric)

		err := grpcClient.InsertItem(context.Background(), &proto.InsertItemRequest{
			TopicName: topic.Name,
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
