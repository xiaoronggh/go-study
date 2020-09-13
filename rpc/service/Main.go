package main

import (
	pb "awesomeProject/rpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

const port = ":50001"

type server struct {

}

func (*server) Upper(ctx context.Context, request *pb.UpperRequest) (*pb.UpperReply, error) {
	log.Printf("received: %s", request.Name)
	return &pb.UpperReply{Message: strings.ToUpper(request.Name)}, nil;
}

func main() {
	listener, err := net.Listen("tcp", port);
	if err != nil {
		log.Fatalf("failed listen port: %v", err)
	}

	serv := grpc.NewServer();
	pb.RegisterToUpperServer(serv, &server{});
	//reflection.Register(serv);

	if err:= serv.Serve(listener); err != nil {
		log.Fatalf("server start failed %v", err)
	}
	log.Println("server start ...");
}
