package main

import (
	"errors"
	"log"
	"main/algos"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var upgrader = websocket.Upgrader{}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/ws", func(c echo.Context) error {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}

		log.Println("Connected!")

		// handleClientMessages(ws)
		algos.Start(ws)

		return nil
	})

	e.Logger.Fatal(e.Start(":9090"))
}

func handleClientMessages(ws *websocket.Conn) {
	defer ws.Close()

	for {
		// Read message from the client
		// _, message, err := ws.ReadMessage()
		// if err != nil {
		// 	log.Println("Failed to read message from client:", err)
		// 	break
		// }

		// log.Println("Received message:", string(message))

		// Echo the message back to the client
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Println("Failed to send message to client:", err)
			break
		}
	}

	log.Println("Client connection closed")
}
