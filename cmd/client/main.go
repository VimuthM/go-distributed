// Dummy client for testing the server
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "distributed/api/runner"
)

func main() {
	// Set up a connection to the server.
	// Dummy creds for testing
	// TODO: Replace with TLS creds
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail: %v", err)
	}
	log.Println("Client is running at port 50051")
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