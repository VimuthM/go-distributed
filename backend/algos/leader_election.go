package algos

import (
	"fmt"
	"math/rand"
	"time"

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

type StateMessage struct {
	Sender_id   int    `json:"sender_id"`
	Receiver_id int    `json:"receiver_id"`
	Action      string `json:"action"`
}

type InitialStateMessage struct {
	Nodes []int   `json:"nodes"`
	Links [][]int `json:"links"`
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
	master_chans := make([]chan StateMessage, n)

	for i := 0; i < n; i++ {
		channels[i] = make(chan Message)
		master_chans[i] = make(chan StateMessage)
	}

	for i := 0; i < n; i++ {
		nodes[i].id = i
		nodes[i].state = "active"
		nodes[i].round = 0
		nodes[i].leader = -1
	}

	// Randomize
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(n, func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	// Construct the initial state message
	var nodes_arr []int
	var links_arr [][]int

	// Writing to the left, Reading from the right
	for i := 0; i < n; i++ {
		var left int
		if i == 0 {
			left = n - 1
		} else {
			left = i - 1
		}

		right := i

		fmt.Println("Node ", i, " left: ", left, " right: ", right)
		nodes_arr = append(nodes_arr, indices[i])
		// [a, b] : a sends to b
		links_arr = append(links_arr, []int{indices[i], indices[left]})

		go node(indices[i], channels[left], channels[right], done, left, right, master_chans[i])
	}

	// Send the initial state message
	initial_state_msg := InitialStateMessage{
		Nodes: nodes_arr,
		Links: links_arr,
	}
	err := ws.WriteJSON(initial_state_msg)
	if err != nil {
		fmt.Println("Error writing to websocket")
	}

	// Listen to all master channels, on any receive write it to the websocket
	go func() {
		for {
			for _, master_chan := range master_chans {
				select {
				case msg := <-master_chan:
					err := ws.WriteJSON(msg)
					if err != nil {
						fmt.Println("Error writing to websocket")
					}
					// time sleep 1 second
					time.Sleep(1 * time.Second)
				default:
					// fmt.Println("No message received on master channel")
				}
			}
		}
	}()
}

// Message passing clockwise => left neighbour is sent to
func node(id int, write chan<- Message, read <-chan Message, done chan bool, left int, right int, master_chan chan<- StateMessage) {
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
		master_chan <- StateMessage{
			Sender_id:   id,
			Receiver_id: left,
			Action:      "forward",
		}
	}

	for {
		// fmt.Println("Node ", id, " waiting for message")
		select {
		case msg := <-read:
			// print := "Node ", id, " received message from ", msg.sender_id
			print := fmt.Sprintf("Node %d received message from %d", id, msg.sender_id)
			fmt.Println(print)
			if msg.message == id {
				fmt.Println("Node ", id, " is the leader")
				master_chan <- StateMessage{
					Sender_id:   id,
					Receiver_id: left,
					Action:      "leader",
				}
				done <- true
				return
			} else if msg.message > id {
				msg.round++
				write <- msg
				master_chan <- StateMessage{
					Sender_id:   id,
					Receiver_id: left,
					Action:      "forward",
				}
			} else {
				print := fmt.Sprintf("Node %d dropped message %d from %d", id, msg.message, msg.sender_id)
				fmt.Println(print)
				master_chan <- StateMessage{
					Sender_id:   id,
					Receiver_id: left,
					Action:      "drop",
				}
				if !sent {
					write <- id_msg
					master_chan <- StateMessage{
						Sender_id:   id,
						Receiver_id: left,
						Action:      "forward",
					}
					sent = true
				}
			}
		case <-done:
			return
		}
	}
}
