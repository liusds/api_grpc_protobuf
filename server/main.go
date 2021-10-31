package main

import (
	"api_grpc_protobuf/proto"
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterCreateServerServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listen); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, req *proto.Request) (res *proto.Response, err error) {
	a, b := req.GetA(), req.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.Request) (res *proto.Response, err error) {
	a, b := req.GetA(), req.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
