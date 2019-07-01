package database_client

import (
	"context"
	"time"

	"github.com/esvm/microservices/src/proto"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type grpcDatabaseClient struct {
	client proto.DcmApiDatabaseServiceClient
}

func NewGrpcDatabaseClient(addr string) (*grpcDatabaseClient, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinearWithJitter(20*time.Millisecond, 0.5)),
		grpc_retry.WithCodes(codes.Unavailable, codes.Aborted, codes.ResourceExhausted),
	}

	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		return nil, err
	}

	return &grpcDatabaseClient{
		client: proto.NewDcmApiDatabaseServiceClient(conn),
	}, nil
}

func (g *grpcDatabaseClient) InsertItem(ctx context.Context, req *proto.InsertItemRequest) error {
	_, err := g.client.InsertItem(ctx, req)
	return err
}

func (g *grpcDatabaseClient) GetItems(ctx context.Context, req *proto.GetItemsRequest) (*proto.GetItemsResponse, error) {
	return g.client.GetItems(ctx, req)
}
