package algos

import (
	"fmt"

	"github.com/gorilla/websocket"
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

func StartLeaderElection(ws *websocket.Conn, n int) {
	chang_roberts(ws, n)
}

func chang_roberts(ws *websocket.Conn, n int) {
	fmt.Println("Chang and Roberts")
	fmt.Println("n =", n)

	channels := make([]chan Message, n)
	nodes := make([]Node, n)
	done := make(chan bool)
	master_chans := make([]chan string, n)

	for i := 0; i < n; i++ {
		channels[i] = make(chan Message)
		master_chans[i] = make(chan string)
	}

	for i := 0; i < n; i++ {
		nodes[i].id = i
		nodes[i].state = "active"
		nodes[i].round = 0
		nodes[i].leader = -1
	}

	for i := 0; i < n; i++ {
		var left int
		if i == 0 {
			left = n - 1
		} else {
			left = i - 1
		}

		right := i

		fmt.Println("Node ", i, " left: ", left, " right: ", right)
		go node(i, channels[left], channels[right], done, left, right, master_chans[i])
	}

	// Listen to all master channels, on any receive write it to the websocket
	go func() {
		for {
			for _, master_chan := range master_chans {
				select {
				case msg := <-master_chan:
					fmt.Println("Received message on master channel")
					err := ws.WriteMessage(websocket.TextMessage, []byte(msg))
					if err != nil {
						fmt.Println("Error writing to websocket")
					}
				default:
					// fmt.Println("No message received on master channel")
				}
			}
		}
	}()
}

// Message passing clockwise => left neighbour is sent to
func node(id int, write chan<- Message, read <-chan Message, done chan bool, left int, right int, master_chan chan<- string) {
	fmt.Println("Node ", id, " started")
	var id_msg = Message{
		sender_id:   id,
		receiver_id: left,
		round:       0,
		message:     id,
	}
	sent := false
	if id == 0 {
		fmt.Println("Node ", id, " sent message to ", left)
		write <- id_msg
		sent = true
	}

	for {
		// fmt.Println("Node ", id, " waiting for message")
		select {
		case msg := <-read:
			// print := "Node ", id, " received message from ", msg.sender_id
			print := fmt.Sprintf("Node %d received message from %d", id, msg.sender_id)
			fmt.Println(print)
			master_chan <- print
			if msg.message == id {
				fmt.Println("Node ", id, " is the leader")
				done <- true
				return
			} else if msg.message > id {
				msg.round++
				write <- msg
			} else {
				print := fmt.Sprintf("Node %d dropped message %d from %d", id, msg.message, msg.sender_id)
				fmt.Println(print)
				master_chan <- print
				if !sent {
					write <- id_msg
					sent = true
				}
			}
		case <-done:
			return
		}
	}
}
