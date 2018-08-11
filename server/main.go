package main

import (
	"context"
	"grpc-server/proto"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const(
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){

	fmt.Println(req.Name)

	return &simple.HelloReplay{Message:"hello =======> " + req.Name},nil
}

func main(){
	lis,err := net.Listen("tcp",port)

	if err != nil {
		log.Fatal("fail to listen")
	}

	s := grpc.NewServer()

	simple.RegisterSimpleServer(s,&server{})

	reflection.Register(s)

	if err:= s.Serve(lis);err != nil{
		log.Fatal("fail to server")
	}
}