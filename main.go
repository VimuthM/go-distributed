package main

import (
	"fmt"
)

type Node struct {
	id     int
	state  string
	round  int
	leader int
}

type Message struct {
	sender_id   int
	receiver_id int
	round       int
	message     int
}

func main() {

	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()
	// fmt.Println("Server is running at port 9000")

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }

	chang_roberts(10)
}

func chang_roberts(n int) {
	fmt.Println("Chang and Roberts")
	fmt.Println("n = ", n)

	channels := make([]chan Message, n-1)
	nodes := make([]Node, n)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		nodes[i].id = i
		nodes[i].state = "active"
		nodes[i].round = 0
		nodes[i].leader = -1
	}

	for i := 0; i < n; i++ {
		var left int
		if i == 0 {
			left = n - 2
		} else {
			left = i - 1
		}

		var right int
		if i == n-1 {
			right = 0
		} else {
			right = i + 1
		}

		go node(i, channels[left], channels[right], done, left, right)
	}
}

// Message passing clockwise => left neighbour is sent to
func node(id int, write chan<- Message, read <-chan Message, done chan bool, left int, right int) {
	var msg = Message{
		sender_id:   id,
		receiver_id: left,
		round:       0,
		message:     id,
	}
	write <- msg

	for {
		select {
		case msg = <-read:
			fmt.Println("Node ", id, " received message from ", msg.sender_id)
			if msg.message == id {
				fmt.Println("Node ", id, " is the leader")
				done <- true
				return
			} else if msg.message > id {
				msg.round++
				write <- msg
			} else {
				fmt.Println("Dropped message", msg.message, " from ", msg.sender_id, " ", msg.round)
			}
		case <-done:
			return
		}
	}
}
