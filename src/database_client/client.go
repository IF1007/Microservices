package database_client

import (
	"context"

	"github.com/esvm/microservices/src/proto"
)

type DatabaseClient interface {
	InsertItem(context.Context, *proto.InsertItemRequest) error
	GetItems(context.Context, *proto.GetItemsRequest) (*proto.GetItemsResponse, error)
}
