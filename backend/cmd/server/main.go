// Package main implements a sever for runner service
package main

import (
	"fmt"
	"log"
	"net"

	pb "distributed/api/runner"

	"google.golang.org/grpc"
)

const (
	listenAddress = "localhost:9090"
)

type server struct {
	pb.UnimplementedRunnerServer
}

func (*server) RunAlgo(req *pb.RunRequest, stream pb.Runner_RunAlgoServer) error {
	fmt.Printf("Run Algo called: Received: %v\n", req)
	var i int32 = 0
	for i = 0; i < 5; i++ {
		if err := stream.Send(&pb.NodeData{SenderId: i, ReceiverId: i, Round: i, Message: i}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server is running at ", listenAddress)

	s := grpc.NewServer()
	pb.RegisterRunnerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
