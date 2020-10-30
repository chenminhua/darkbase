package main

import (
	"context"
	pb "github.com/chenminhua/darkbase/darkbase"
	"github.com/chenminhua/darkbase/server/conf"
	"github.com/chenminhua/darkbase/server/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":12345"
)


type Server struct {
	pb.UnimplementedDarkbaseServiceServer
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	tableName := req.TableName
	rdb := db.OpenRocksdb(tableName)
	b, _ := rdb.Get(req.Key)
	log.Printf("get key %v, with response %v", req.Key, b)
	return &pb.GetResponse{ErrCode: pb.ErrorCode_Success, Value: b}, nil
}

func (s *Server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	//log.Printf("put key %v = %v", req.Key, req.Value)
	tableName := req.TableName
	rdb := db.OpenRocksdb(tableName)
	err := rdb.Put(req.Key, req.Value)
	if err != nil {
		log.Printf("put key err %v", err)
	}
	return &pb.PutResponse{ErrCode: pb.ErrorCode_Success}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDarkbaseServiceServer(s, &Server{})
	conf.SetBatchConfig(true)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
