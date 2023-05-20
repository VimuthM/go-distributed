// Dummy client for testing the server
package main

import (
	"context"
	"log"

	pb "distributed/api/runner"

	"google.golang.org/grpc"
)

const listenAddress = "localhost:9090"

func main() {
	// Set up a connection to the server.
	// Dummy creds for testing
	// TODO: Replace with TLS creds
	conn, err := grpc.Dial(listenAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail: %v", err)
	}
	log.Println("Client is running at ", listenAddress)
	defer conn.Close()

	c := pb.NewRunnerClient(conn)

	// Call the server to run the algorithm
	stream, err := c.RunAlgo(context.Background(), &pb.RunRequest{Name: "chang_roberts"})
	if err != nil {
		log.Fatalf("Failed to get data stream: %v", err)
	}

	// Receive data from stream
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive data: %v", err)
		}

		log.Printf("Response from server: %v", response)
	}
}
