package main

import (
	pb "awesomeProject/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
)

const address = "127.0.0.1:50001"

//func main() {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		log.Fatal("grpc dial failed: %v", err);
//	}
//
//	defer conn.Close();
//
//	client := pb.NewToUpperClient(conn);
//
//	name := "hello"
//	if len(os.Args) > 1 {
//		name = os.Args[1]
//	}
//
//	reply, err := client.Upper(context.Background(), &pb.UpperRequest{Name: name})
//	if err != nil {
//		log.Fatalf("upper failed %v", err)
//	}
//	log.Printf("%s upper is %s ", name, reply.Message)
//}

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc dail failed %v", err);
	}

	defer conn.Close()
	client := pb.NewToUpperClient(conn);

	name := "hello";
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	reply, err := client.Upper(context.Background(), &pb.UpperRequest{Name: name});
	if err != nil {
		log.Fatalf("upper err %v", err);
	}
	log.Printf("%s upper is %s", name, reply.Message)

}
