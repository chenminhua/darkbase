package client

import (
	"context"
	pb "github.com/chenminhua/darkbase/darkbase"
	"google.golang.org/grpc"
	"log"
	"time"
)

type DarkbaseClient struct {
	conn *grpc.ClientConn
	client pb.DarkbaseServiceClient
}

func NewClient(address string) (*DarkbaseClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := pb.NewDarkbaseServiceClient(conn)
	return &DarkbaseClient{conn: conn, client: c}, nil
}

func (client *DarkbaseClient) Put(tablename string, key string, value []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.client.Put(ctx, &pb.PutRequest{Key: key, Value: value, TableName: tablename})
	return err
}

func (client *DarkbaseClient) Get(tablename string, key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.client.Get(ctx, &pb.GetRequest{Key: key, TableName: tablename})
	return r.Value, err
}

func (client *DarkbaseClient) BatchPut(tablename string, putReqs []*pb.PutRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var req = &pb.BatchPutRequest{TableName: tablename, Req: putReqs}
	_, err := client.client.BatchPut(ctx, req)
	return err
}

func (client *DarkbaseClient) BatchGet(tablename string, getReqs []*pb.GetRequest)

func (client *DarkbaseClient) Close() error{
	return client.conn.Close()
}
