package main

import (
	"encoding/json"
	"errors"
	"log"
	"main/algos"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var upgrader = websocket.Upgrader{}

type ClientMessage struct {
	Algorithm string `json:"algorithm"`
	Nodes     int    `json:"nodes"`
}

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

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("Failed to read message from client:", err)
				return nil
			}

			var cmsg ClientMessage
			err = json.Unmarshal(message, &cmsg)
			if err != nil {
				log.Println("Failed to unmarshal message:", err)
				return nil
			}

			if cmsg.Algorithm == "chang_roberts" {
				algos.StartLeaderElection(ws, cmsg.Nodes)
			}
		}
	})

	e.Logger.Fatal(e.Start(":9090"))
}
